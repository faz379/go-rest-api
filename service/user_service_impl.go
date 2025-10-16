package service

import (
	"belajar-rest-api-golang/helper"
	"belajar-rest-api-golang/model/domain"
	"belajar-rest-api-golang/model/web"
	"belajar-rest-api-golang/repository"
	"belajar-rest-api-golang/util"
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewUserService(userRepository repository.UserRepository, db *sql.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             db,
		Validate:       validate,
	}
}

func (service *UserServiceImpl) Login(ctx context.Context, tx *sql.Tx, request web.UserLoginRequest) (string, error) {
	if request.Email == "" {
		return "", fmt.Errorf("email tidak boleh kosong")
	}
	if request.Password == "" {
		return "", fmt.Errorf("password tidak boleh kosong")
	}

	// Jika tx nil, buka transaction baru
	var err error
	if tx == nil {
		tx, err = service.DB.BeginTx(ctx, nil)
		if err != nil {
			return "", err
		}
		defer helper.CommitOrRollback(tx)
	}

	user, err := service.UserRepository.FindByEmail(ctx, tx, request.Email)
	if err != nil {
		return "", fmt.Errorf("user tidak ditemukan")
	}

	if !util.CheckPassword(user.Password, request.Password) {
		return "", fmt.Errorf("password salah")
	}

	token, err := util.GenerateToken(user.Id, time.Hour*24)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (service *UserServiceImpl) Register(ctx context.Context, tx *sql.Tx, request web.UserRegisterRequest) (web.UserResponse, error) {
	// Validasi input
	if request.Username == "" {
		return web.UserResponse{}, fmt.Errorf("username tidak boleh kosong")
	}
	if request.Email == "" {
		return web.UserResponse{}, fmt.Errorf("email tidak boleh kosong")
	}
	if request.Password == "" {
		return web.UserResponse{}, fmt.Errorf("password tidak boleh kosong")
	}

	if tx == nil {
		var err error
		tx, err = service.DB.BeginTx(ctx, nil)
		if err != nil {
			return web.UserResponse{}, err
		}
		defer helper.CommitOrRollback(tx) // pakai helper di sini
	}

	existingUser, _ := service.UserRepository.FindByEmail(ctx, tx, request.Email)
	if existingUser.Id != 0 {
		return web.UserResponse{}, fmt.Errorf("email sudah digunakan")
	}

	hashed, err := util.HashPassword(request.Password)
	if err != nil {
		return web.UserResponse{}, err
	}

	user := domain.User{
		Username: request.Username,
		Email:    request.Email,
		Password: hashed,
	}

	saved := service.UserRepository.Save(ctx, tx, user)

	response := web.UserResponse{
		Id:       saved.Id,
		Username: saved.Username,
	}

	return response, nil
}
