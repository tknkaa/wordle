package api

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func SetupRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("logged")
		w.Write([]byte("Hello, world!\n"))
	})
	return r
}

func VercelHandler(w http.ResponseWriter, r *http.Request) {
	router := SetupRouter()
	router.ServeHTTP(w, r)
}
