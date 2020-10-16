package article

import "time"

//go:generate fixtory -type=Author -output=author.fixtory.go

// Author represents article's author
type Author struct {
	ID   int
	Name string
}

// Article represents article
type Article struct {
	ID                 int
	Title              string
	Body               string
	AuthorID           int
	PublishScheduledAt time.Time
	PublishedAt        time.Time
	Status             ArticleStatus
	LikeCount          int
}

// ArticleStatus represents ArticleStatus
type ArticleStatus int
