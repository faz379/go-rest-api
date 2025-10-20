package domain

type Comment struct {
	Id       int
	PostId   int
	Content  string
	AuthorId int
}
