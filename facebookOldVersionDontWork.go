package crawlers

// facebook hashtag is now only for logged in users

/* import (
	"fmt"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

func getPostsFromFacebook(urlToCrawl string) []Post {
	postsToReturn := []Post{}

	c := colly.NewCollector()

	c.OnResponse(func(r *colly.Response) {
		htmlString := string(r.Body)
		splittedInfo := strings.Split(htmlString, "<abbr title=")
		if len(splittedInfo) > 1 {
			for i := 1; i < len(splittedInfo); i++ {
				splitted := splittedInfo[i]
				datum := splitted[1:len(splitted)]
				datum = datum[0:strings.Index(datum, "\"")]
				if strings.Index(splitted, "<p>") > -1 {
					splitted = splitted[strings.Index(splitted, "<p>")+len("<p>") : len(splitted)]
					if strings.Index(splitted, "</p>") > 0 {
						splitted = splitted[0:strings.Index(splitted, "</p>")]
						for strings.Index(splitted, "<") > -1 {
							if strings.Index(splitted, ">") > -1 {
								if strings.Index(splitted, "<") < strings.Index(splitted, ">") {
									splittedFistPart := splitted[0:strings.Index(splitted, "<")]
									splittedSecondPart := splitted[strings.Index(splitted, ">")+1 : len(splitted)]
									splitted = splittedFistPart + splittedSecondPart
								} else {
									splitted = splitted[strings.Index(splitted, ">")+1 : len(splitted)]
								}
							} else {
								splitted = splitted[0:strings.Index(splitted, "<")]
							}
						}
						if len(splitted) > MaxTitleLenght {
							splitted = splitted[0:MaxTitleLenght] + "..."
						}
						splitted = FilterNewLines(splitted)
						postToReturn := Post{Date: datum, Title: splitted, Link: urlToCrawl}
						postsToReturn = append(postsToReturn, postToReturn)
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

// FilterNewLines change bit code new lines
func FilterNewLines(s string) string {
	return strings.Map(func(r rune) rune {
		switch r {
		case 0x000A, 0x000B, 0x000C, 0x000D, 0x0085, 0x2028, 0x2029:
			return -1
		default:
			return r
		}
	}, s)
}

// FindFacebookInfo return posts
func FindFacebookInfo(facebookToCrawl string) SocialNetwork {
	facebookToReturn := SocialNetwork{}

	if strings.Contains(facebookToCrawl, " ") {
		facebookToCrawl = strings.Replace(facebookToCrawl, " ", "", -1)
	}

	networkURLToCrawl := "https://de-de.facebook.com/hashtag/" + facebookToCrawl
	posts := getPostsFromFacebook(networkURLToCrawl)
	facebookToReturn.Tag = facebookToCrawl
	facebookToReturn.Count = len(posts)
	facebookToReturn.Posts = posts
	if len(posts) > 0 {
		layout := "02.01.2006 15:04"
		newestDate := posts[0].Date
		newestDateTime, err := time.Parse(layout, newestDate)
		if err != nil {
			fmt.Println(err)
		}

		for _, post := range posts {
			thisPostDate, err := time.Parse(layout, post.Date)
			if err != nil {
				fmt.Println(err)
			} else {
				if thisPostDate.After(newestDateTime) {
					newestDate = post.Date
					newestDateTime = thisPostDate
				}
			}
		}
		facebookToReturn.NewestPostDate = newestDate
	}

	return facebookToReturn
}
*/
