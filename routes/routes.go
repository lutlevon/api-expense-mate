package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Routes() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hi from chi"))
	})
	return r
}
