package handlers

import (
	"context"
	"project/internal/userService"
	"project/internal/web/users"
)

type UserHandler struct {
	Service *userService.UserService
}

func NewUserHandler(service *userService.UserService) *UserHandler {
	return &UserHandler{Service: service}
}

func (u *UserHandler) GetUsers(context.Context, users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	allUsers, err := u.Service.GetUsers()
	if err != nil {
		return nil, err
	}

	response := users.GetUsers200JSONResponse{}
	for _, usr := range allUsers {
		user := users.User{
			Id:    &usr.ID,
			Email: &usr.Email,
		}
		response = append(response, user)
	}

	return response, nil
}

func (u *UserHandler) PostUsers(_ context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	userRequest := request.Body
	createdUser, err := u.Service.CreateUser(*userRequest.Email, *userRequest.Password)
	if err != nil {
		return nil, err
	}

	response := users.PostUsers201JSONResponse{
		Id:    &createdUser.ID,
		Email: &createdUser.Email,
	}

	return response, nil
}

func (u *UserHandler) PatchUsersId(_ context.Context, request users.PatchUsersIdRequestObject) (users.PatchUsersIdResponseObject, error) {
	userRequest := request.Body
	err := u.Service.UpdateUser(request.Id, *userRequest.Email, *userRequest.Password)
	if err != nil {
		return nil, err
	}

	return users.PatchUsersId200JSONResponse{}, nil
}

func (u *UserHandler) DeleteUsersId(_ context.Context, request users.DeleteUsersIdRequestObject) (users.DeleteUsersIdResponseObject, error) {
	err := u.Service.DeleteUser(request.Id)
	if err != nil {
		return nil, err
	}

	return users.DeleteUsersId204Response{}, nil
}
