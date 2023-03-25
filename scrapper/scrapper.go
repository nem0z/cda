package scrapper

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/nem0z/cda/types"
)

const baseUrl = "https://cda.chronomania.net"
const path = "/forum_entry.php"

type Processed struct {
	index int
	post  *types.Post
}

type Rejected struct {
	err   error
	index int
}

func worker(indexs <-chan int, result chan *Processed, reject chan *Rejected) {
	for i := range indexs {
		res, err := ProcessOne(i)

		if err != nil || res == nil || res.Content == "" {
			reject <- &Rejected{err, i}
			continue
		}
		result <- &Processed{i, res}
	}
}

func Process(start int, n int) map[int]*types.Post {
	processing := true
	results := make(map[int]*types.Post)

	jobs := make(chan int, n)
	result := make(chan *Processed, n)
	reject := make(chan *Rejected)

	for i := 0; i < 25; i++ {
		go worker(jobs, result, reject)
	}

	go func() {
		for processing {
			rej := <-reject
			if _, ok := rej.err.(DoNotExistError); ok {
				fmt.Printf("%v doesn't exist\n", rej.index)
				result <- &Processed{rej.index, nil}
			}
			jobs <- rej.index
		}
	}()

	for i := start; i >= start-n; i-- {
		jobs <- i
	}

	for i := 0; i <= n; i++ {
		select {
		case processed := <-result:
			results[processed.index] = processed.post
		}
	}

	return results
}

func ProcessOne(index int) (*types.Post, error) {
	url := fmt.Sprintf("%v%v?id=%v", baseUrl, path, index)

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, StatusError{res}
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	if doc.Find("title").First().Text() == "Coin des Affaires" {
		return nil, DoNotExistError{}
	}

	headlineText := doc.Find(".postingheadline").First().Text()
	authorSection := doc.Find(".author").First()
	post := doc.Find(".posting").First()
	images := ParseImages(post)
	signature, err := doc.Find(".signature").First().Html()
	if err != nil {
		return nil, err
	}

	headline := ParseHeadline(headlineText)
	author := ParseAuthor(authorSection)

	return &types.Post{
		Index:     index,
		Head:      headline,
		Author:    author,
		Content:   fmt.Sprintf("%q\n", post.Text()),
		Images:    images,
		Signature: fmt.Sprintf("%q\n", signature),
	}, nil
}
