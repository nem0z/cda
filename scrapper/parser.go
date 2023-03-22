package scrapper

import (
	"regexp"

	"github.com/PuerkitoBio/goquery"
)

func ParseHeadline(headline string) *Headline {
	re := regexp.MustCompile(`^\[(.*?)\]\s(.*?)\s-\s(\d+)\&euro;$`)
	matches := re.FindStringSubmatch(headline)

	if len(matches) == 4 {
		return &Headline{
			matches[1],
			matches[2],
			matches[3],
		}
	}
	return nil
}

func ParseAuthor(div *goquery.Selection) *Author {
	return &Author{
		div.Find("p.author b").Text(),
		div.Find("p.author").Contents().Last().Text(),
	}
}

func ParseImages(post *goquery.Selection) (images []string) {
	post.Find("img").Each(func(i int, element *goquery.Selection) {
		src, ok := element.Attr("src")
		if ok {
			images = append(images, src)
		}
	})
	return images
}
