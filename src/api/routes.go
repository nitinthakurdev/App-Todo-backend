package api

import (
	"net/http"

	"github.com/nitinthakurdev/todo-app-backend/src/controller"
)

func (app *Application) Mount() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", controller.Health)

	// user routes
	mux.HandleFunc("POST /api/v1/user/create", controller.CreateUser)
	mux.HandleFunc("POST /api/v1/user/login", controller.LoginUser)

	return mux
}
