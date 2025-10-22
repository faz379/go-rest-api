package domain

import "time"

type Post struct {
	Id        int
	Title     string
	Content   string
	ImageURL  string
	AuthorId  int
	CreatedAt time.Time
}
