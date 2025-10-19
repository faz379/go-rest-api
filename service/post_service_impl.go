package service

import (
	"belajar-rest-api-golang/exception"
	"belajar-rest-api-golang/helper"
	"belajar-rest-api-golang/model/domain"
	"belajar-rest-api-golang/model/web"
	"belajar-rest-api-golang/repository"
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
)

type PostServiceImpl struct {
	PostRepository repository.PostRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewPostService(postRepository repository.PostRepository, DB *sql.DB, validate *validator.Validate) PostService {
	return &PostServiceImpl{
		PostRepository: postRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (service *PostServiceImpl) Create(ctx context.Context, request web.PostCreateRequest, userId int) web.PostResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	post := domain.Post{
		Id:       0,
		Title:    request.Title,
		Content:  request.Content,
		ImageURL: request.ImageURL,
		AuthorId: userId,
	}

	post = service.PostRepository.Save(ctx, tx, post)

	return helper.ToPostResponse(post)
}
func (service *PostServiceImpl) Update(ctx context.Context, request web.PostUpdateRequest) web.PostResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	post, err := service.PostRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	post.Title = request.Title
	post.Content = request.Content
	post.ImageURL = request.ImageURL

	post = service.PostRepository.Update(ctx, tx, post)

	return helper.ToPostResponse(post)
}

func (service *PostServiceImpl) Delete(ctx context.Context, postId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	post, err := service.PostRepository.FindById(ctx, tx, postId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.PostRepository.Delete(ctx, tx, post)
}

func (service *PostServiceImpl) FindById(ctx context.Context, postId int) web.PostResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	post, err := service.PostRepository.FindById(ctx, tx, postId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToPostResponse(post)
}

func (service *PostServiceImpl) FindAll(ctx context.Context) []web.PostResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	posts := service.PostRepository.FindAll(ctx, tx)

	return helper.ToPostResponses(posts)
}
