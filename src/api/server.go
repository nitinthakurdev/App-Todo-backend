package api

import (
	"log"
	"net/http"
	"time"

	"github.com/nitinthakurdev/todo-app-backend/src/config"
	"github.com/nitinthakurdev/todo-app-backend/src/database"
)

type Application struct {
	Addr string
}

func InitServer(Addr string) *Application {
	return &Application{
		Addr: Addr,
	}
}

func (app *Application) StartServer(mux *http.ServeMux) error {

	server := &http.Server{
		Addr:         app.Addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}
	log.Printf("server is up and running on port %s\n", app.Addr)
	startup()
	return server.ListenAndServe()
}

func startup() {
	if err := database.Mongo(config.Keys().DB); err != nil {
		panic(err)
	}
}
