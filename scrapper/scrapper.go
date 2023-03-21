package scrapper

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

const baseUrl = "https://cda.chronomania.net"
const path = "/forum_entry.php"

type Headline struct {
	status string
	title  string
	price  string
}

type Author struct {
	name string
	date string
}

type Post struct {
	head      *Headline
	author    *Author
	Content   string
	imgs      []string
	signature string
}

func Process(index int) (*Post, error) {
	url := fmt.Sprintf("%v%v?id=%v", baseUrl, path, index)

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("status code error: (%d) %s", res.StatusCode, res.Status))
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
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

	return &Post{
		headline,
		author,
		post.Text(),
		images,
		signature,
	}, nil
}
