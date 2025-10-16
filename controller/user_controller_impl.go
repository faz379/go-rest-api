package controller

import (
	"belajar-rest-api-golang/helper"
	"belajar-rest-api-golang/model/web"
	"belajar-rest-api-golang/service"
	"net/http"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) *UserControllerImpl {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (controller *UserControllerImpl) Register(writer http.ResponseWriter, request *http.Request) {
	var registerRequest web.UserRegisterRequest
	helper.ReadFromRequestBody(request, &registerRequest)

	response, err := controller.UserService.Register(request.Context(), nil, registerRequest)
	if err != nil {
		helper.WriteErrorResponse(writer, http.StatusBadRequest, err.Error())
		return
	}

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   response,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) Login(writer http.ResponseWriter, request *http.Request) {
	var loginRequest web.UserLoginRequest
	helper.ReadFromRequestBody(request, &loginRequest)

	token, err := controller.UserService.Login(
		request.Context(),
		nil, // tx bisa nil jika pakai helper CommitOrRollback di service
		loginRequest,
	)
	if err != nil {
		helper.WriteErrorResponse(writer, http.StatusUnauthorized, err.Error())
		return
	}

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data: map[string]string{
			"token": token,
		},
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) Logout(writer http.ResponseWriter, request *http.Request) {
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "LOGOUT SUCCESS",
	}

	helper.WriteToResponseBody(writer, webResponse)
}
