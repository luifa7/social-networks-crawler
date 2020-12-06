package crawlers

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

func getPostsFromInstagram(urlToCrawl string) []Post {
	postsToReturn := []Post{}
	c := colly.NewCollector()

	c.OnResponse(func(r *colly.Response) {
		htmlString := string(r.Body)
		if (strings.Index(htmlString, "window._sharedData =")) > 0 {
			jsonData := htmlString[strings.Index(htmlString, "window._sharedData =")+len("window._sharedData =") : len(htmlString)]
			jsonData = strings.TrimSpace(jsonData[0:strings.Index(jsonData, ";</script>")])

			if strings.Index(jsonData, "\"edge_hashtag_to_media\":") > 0 {
				jsonData = jsonData[strings.Index(jsonData, "\"edge_hashtag_to_media\":")+len("\"edge_hashtag_to_media\":") : len(jsonData)]
				jsonData = strings.TrimSpace(jsonData[0:strings.Index(jsonData, ",\"edge_hashtag_to_top_posts\"")])
				byt := []byte(jsonData)
				var payload map[string]interface{}
				if err := json.Unmarshal(byt, &payload); err != nil {
					panic(err)
				}
				postsList := payload["edges"].([]interface{})

				postsToSave := MaxPostsSize
				if len(postsList) < MaxPostsSize {
					postsToSave = len(postsList)
				}
				for i := 0; i < postsToSave; i++ {
					postJSON := postsList[i].(map[string]interface{})
					node := postJSON["node"].(map[string]interface{})
					takenAt := int64(node["taken_at_timestamp"].(float64))
					datum := time.Unix(takenAt, 0).Format("02.01.2006 15:04:05")
					link := "https://www.instagram.com/p/" + node["shortcode"].(string)
					title := link
					if len(title) > MaxTitleLenght {
						title = title[0:MaxTitleLenght] + "..."
					}
					postToReturn := Post{Date: datum, Title: title, Link: link}
					postsToReturn = append(postsToReturn, postToReturn)
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

// FindInstagramInfo return posts
func FindInstagramInfo(hashtagToCrawl string) SocialNetwork {
	hashtagToCrawl = normalizeWithSymbol(hashtagToCrawl, "")
	urlToCrawl := "https://www.instagram.com/explore/tags/" + hashtagToCrawl
	posts := getPostsFromInstagram(urlToCrawl)

	return makeSocialNetwork("instagram", hashtagToCrawl, posts)
}
