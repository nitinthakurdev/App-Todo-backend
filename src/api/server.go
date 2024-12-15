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
		Handler:      corsMiddleware(mux),
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

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // Allow all origins
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight OPTIONS request
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
