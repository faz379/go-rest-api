package helper

import (
	"belajar-rest-api-golang/model/domain"
	"belajar-rest-api-golang/model/web"
)

func ToPostResponse(post domain.Post) web.PostResponse {
	return web.PostResponse{
		Id:       post.Id,
		Title:    post.Title,
		Content:  post.Content,
		ImageURL: post.ImageURL,
		AuthorId: post.AuthorId,
	}
}

func ToCommentResponse(comment domain.Comment) web.CommentResponse {
	return web.CommentResponse{
		Id:       comment.Id,
		PostId:   comment.PostId,
		Content:  comment.Content,
		AuthorId: comment.AuthorId,
	}
}

func ToPostResponses(posts []domain.Post) []web.PostResponse {
	var postResponses []web.PostResponse
	for _, post := range posts {
		postResponses = append(postResponses, ToPostResponse(post))
	}
	return postResponses
}
