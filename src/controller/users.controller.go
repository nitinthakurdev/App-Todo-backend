package controller

import (
	"fmt"
	"net/http"

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

	result, err := services.CreateUser(user)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	var NewData = &types.UserResponse{
		Username: result.Username,
		Email:    result.Email,
		Message:  "User Created successful",
		Token:    "token",
	}
	utils.WriteJson(w, http.StatusOK, NewData)
}
