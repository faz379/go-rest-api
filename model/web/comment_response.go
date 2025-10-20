package web

type CommentResponse struct {
	Id       int    `json:"id"`
	PostId   int    `json:"post_id"`
	AuthorId int    `json:"author_id"`
	Content  string `json:"content"`
}
