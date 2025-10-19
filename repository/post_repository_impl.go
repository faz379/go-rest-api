package repository

import (
	"belajar-rest-api-golang/helper"
	"belajar-rest-api-golang/model/domain"
	"context"
	"database/sql"
	"errors"
)

type PostRepositoryImpl struct {
}

func NewPostRepository() PostRepository {
	return &PostRepositoryImpl{}
}

func (repository *PostRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, post domain.Post) domain.Post {
	SQL := "insert into posts(title, content, image_url, author_id) values (?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, post.Title, post.Content, post.ImageURL, post.AuthorId)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	post.Id = int(id)
	return post
}

func (repository *PostRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, post domain.Post) domain.Post {
	SQL := "update posts set title = ?, content = ?, image_url = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, post.Title, post.Content, post.ImageURL, post.Id)
	helper.PanicIfError(err)

	return post
}

func (repository *PostRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, post domain.Post) {
	SQL := "delete from posts where id = ?"
	_, err := tx.ExecContext(ctx, SQL, post.Id)
	helper.PanicIfError(err)

}

func (repository *PostRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, postId int) (domain.Post, error) {
	SQL := "Select id, title, content, image_url, author_id from posts where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, postId)
	helper.PanicIfError(err)
	defer rows.Close()

	post := domain.Post{}
	if rows.Next() {
		err := rows.Scan(&post.Id, &post.Title, &post.Content, &post.ImageURL, &post.AuthorId)
		helper.PanicIfError(err)
		return post, nil
	} else {
		return post, errors.New("post is not found")
	}
}

func (repository *PostRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Post {
	SQL := "Select id, title, content, image_url, author_id from posts"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var posts []domain.Post
	for rows.Next() {
		post := domain.Post{}
		err := rows.Scan(&post.Id, &post.Title, &post.Content, &post.ImageURL, &post.AuthorId)
		helper.PanicIfError(err)
		posts = append(posts, post)
	}
	return posts
}
