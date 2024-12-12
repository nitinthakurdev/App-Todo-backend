package main

import (
	"log"

	"github.com/nitinthakurdev/todo-app-backend/src/api"
	"github.com/nitinthakurdev/todo-app-backend/src/config"
)

func main() {
	app := &api.Application{
		Addr: config.Keys().Port,
	}

	mux := app.Mount()

	if err := app.StartServer(mux); err == nil {
		log.Panic("server not started", err)
	}
}
