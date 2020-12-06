package crawlers

// Post is the model for the posts
type Post struct {
	Date  string
	Link  string
	Title string
}

// MaxPostsSize maximal wanted posts
const MaxPostsSize = 10

// MaxTitleLenght maximal title lenght
const MaxTitleLenght = 45
