package crawlers

import (
	"fmt"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

func getPostsFromMedium(urlToCrawl string) []Post {
	postsToReturn := []Post{}

	c := colly.NewCollector()

	c.OnHTML("a.link--darken", func(e *colly.HTMLElement) {
		isPostAttr := e.Attr("data-action")
		if strings.Compare(isPostAttr, "open-post") == 0 {
			link := e.Attr("href")
			title := link
			datum := e.ChildAttr("time", "datetime")
			layout := "2006-01-02T15:04:05.000Z"
			parsedDatum, err := time.Parse(layout, datum)
			if err != nil {
				datum = ""
				fmt.Println(err)
			} else {
				datum = parsedDatum.Format("02.01.2006 15:04:05")
			}
			if len(link) > 0 && strings.Contains(link, "medium.com") {
				if len(title) > MaxTitleLenght {
					title = title[0:MaxTitleLenght] + "..."
				}
				postToReturn := Post{Date: datum, Title: title, Link: link}
				postsToReturn = append(postsToReturn, postToReturn)
			}
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit(urlToCrawl)

	return postsToReturn
}

// FindMediumInfo return posts
func FindMediumInfo(hashtagToCrawl string) SocialNetwork {
	hashtagToCrawl = normalizeWithSymbol(hashtagToCrawl, "-")
	urlToCrawl := "https://medium.com/tag/" + hashtagToCrawl + "/latest"
	posts := getPostsFromMedium(urlToCrawl)

	return makeSocialNetwork("medium", hashtagToCrawl, posts)
}
