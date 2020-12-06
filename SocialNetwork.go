package crawlers

// SocialNetwork is the model for the SN results
type SocialNetwork struct {
	Name           string
	Tag            string
	Count          int
	Posts          []Post
	NewestPostDate string
}

// NewSocialNetwork creates a new struct with all the values received from crawling
// the this social network for the decided hashtag
func makeSocialNetwork(name string, tag string, posts []Post) SocialNetwork {

	if len(posts) > MaxPostsSize {
		posts = posts[0:MaxPostsSize]
	}

	sn := SocialNetwork{name, tag, len(posts), posts, ""}

	if len(posts) > 0 {
		sn.NewestPostDate = posts[0].Date
	}
	return sn
}
