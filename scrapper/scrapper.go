package scrapper

import (
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
		return nil, &StatusError{res}
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	if doc.Find("title").First().Text() == "Coin des Affaires" {
		return nil, &DoNotExistError{}
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
