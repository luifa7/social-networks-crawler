package crawlers

import (
	"fmt"

	"github.com/gocolly/colly"
)

func getPostsFromTwitter(urlToCrawl string) []Post {
	postsToReturn := []Post{}

	c := colly.NewCollector()

	c.OnResponse(func(r *colly.Response) {
		htmlString := string(r.Body)
		// fmt.Println(htmlString)
	})

	/* 	c.OnHTML("small.time a", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		a := e.ChildAttr("span", "data-time")
		i, err := strconv.ParseInt(a, 10, 64)
		if err != nil {
			panic(err)
		}
		datum := time.Unix(i, 0).Format("02.01.2006 15:04:05")
		if len(link) > 0 {
			link = "https://twitter.com" + link
			title := link
			if len(title) > MaxTitleLenght {
				title = title[0:MaxTitleLenght] + "..."
			}
			postToReturn := Post{Date: datum, Title: title, Link: link}
			postsToReturn = append(postsToReturn, postToReturn)
		}
	}) */

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit(urlToCrawl)

	return postsToReturn
}

// FindTwitterInfo return posts
func FindTwitterInfo(hashtagToCrawl string) SocialNetwork {
	hashtagToCrawl = normalizeWithSymbol(hashtagToCrawl, "")
	urlToCrawl := "https://twitter.com/search?q=%23" + hashtagToCrawl + "?vertical=default"

	posts := getPostsFromTwitter(urlToCrawl)

	return makeSocialNetwork("twitter", hashtagToCrawl, posts)
}
