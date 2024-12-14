package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/nitinthakurdev/todo-app-backend/src/models"
	"github.com/nitinthakurdev/todo-app-backend/src/services"
	"github.com/nitinthakurdev/todo-app-backend/src/types"
	"github.com/nitinthakurdev/todo-app-backend/src/utils"
	"github.com/nitinthakurdev/todo-app-backend/src/validations"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var user *models.UserModel

	if err := utils.ParseJson(r, &user); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	data, _ := services.FindUser(user)
	if data != nil {
		utils.WriteError(w, http.StatusAlreadyReported, fmt.Errorf("user already exist"))
		return
	}

	if err := validations.CheckValidation(user); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	hash, _ := utils.HashPassword(user.Password)
	var newUser = &models.UserModel{
		Email:    user.Email,
		Username: user.Username,
		Password: hash,
	}

	result, err := services.CreateUser(newUser)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	token, _ := utils.SignToken(result.Email)
	var NewData = &types.UserResponse{
		Username: result.Username,
		Email:    result.Email,
		Message:  "User Created successful",
		Token:    token,
	}
	utils.WriteJson(w, http.StatusOK, NewData)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var user *models.UserModel

	if err := utils.ParseJson(r, &user); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	data, err := services.FindUser(user)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("user not exist"))
		return
	}

	compare := utils.ComparePassword(user.Password, data.Password)
	if !compare {
		utils.WriteError(w, http.StatusBadGateway, fmt.Errorf("Wrong password"))
		return
	}

	token, _ := utils.SignToken(data.Email)

	var NewData = &types.UserResponse{
		Username: data.Username,
		Email:    data.Email,
		Message:  "Login successful",
		Token:    token,
	}

	cookies := &http.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
	}

	http.SetCookie(w, cookies)

	utils.VerifyToken(token)

	utils.WriteJson(w, http.StatusOK, NewData)
}
