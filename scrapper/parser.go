package scrapper

import (
	"regexp"

	"github.com/PuerkitoBio/goquery"
	"github.com/nem0z/cda/types"
)

func ParseHeadline(headline string) *types.Headline {
	re := regexp.MustCompile(`^\[(.*?)\]\s(.*?)\s-\s(\d+)\&euro;$`)
	matches := re.FindStringSubmatch(headline)

	if len(matches) == 4 {
		return &types.Headline{
			Status: matches[1],
			Title:  matches[2],
			Price:  matches[3],
		}
	}
	return nil
}

func ParseAuthor(div *goquery.Selection) *types.Author {
	return &types.Author{
		Name: div.Find("p.author b").Text(),
		Date: div.Find("p.author").Contents().Last().Text(),
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
