package crawlers

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

func getPostsFromYoutube(urlToCrawl string) []Post {
	postsToReturn := []Post{}

	c := colly.NewCollector()

	c.OnResponse(func(r *colly.Response) {
		htmlString := string(r.Body)
		somethingFound := true
		if strings.Contains(htmlString, "Es werden Ergebnisse angezeigt") && strings.Contains(htmlString, "Keine Ergebnisse gefunden") {
			somethingFound = false
		} else if strings.Contains(htmlString, "Showing results for") && strings.Contains(htmlString, "No results found for") {
			somethingFound = false
		}
		if somethingFound {
			start := strings.Index(htmlString, "{\"contents\":[{\"videoRenderer\":")
			end := strings.Index(htmlString, "{\"searchSubMenuRenderer\"")
			if start > 0 && end > 0 {
				substringToWorkWith := htmlString[start:end]

				splittedInfo := strings.Split(substringToWorkWith, "{\"videoRenderer\":")
				if len(splittedInfo) > 1 {
					for i := 1; i < len(splittedInfo); i++ {
						splitted := splittedInfo[i]

						start = strings.Index(splitted, "\"videoId\":\"") + len("\"videoId\":\"")
						end = strings.Index(splitted, "\",\"")
						if start > 0 && end > 0 {
							link := "https://www.youtube.com/watch?v=" + splitted[start:end]

							start = strings.Index(splitted, "{\"text\":\"") + len("{\"text\":\"")
							end = strings.Index(splitted, "\"}]")
							if start > 0 && end > 0 {
								title := splitted[start:end]

								start = strings.Index(splitted, "\"publishedTimeText\":{\"simpleText\":\"") + len("\"publishedTimeText\":{\"simpleText\":\"")
								end = strings.Index(splitted, "\"},\"lengthText\"")
								if start > 0 && end > 0 {
									date := splitted[start:end]

									postToReturn := Post{Date: date, Title: title, Link: link}
									postsToReturn = append(postsToReturn, postToReturn)
								}
							}
						}
					}
				}
			}

		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit(urlToCrawl)

	return postsToReturn
}

// FindYoutubeInfo return posts
func FindYoutubeInfo(hashtagToCrawl string) SocialNetwork {
	hashtagToCrawl = normalizeWithSymbol(hashtagToCrawl, "")
	urlToCrawl := "https://www.youtube.com/results?sp=CAJCBAgBEgA%253D&search_query=%23" + hashtagToCrawl
	posts := getPostsFromYoutube(urlToCrawl)

	return makeSocialNetwork("youtube", hashtagToCrawl, posts)
}
