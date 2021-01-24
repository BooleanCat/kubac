package router

import (
	"net/http"

	"github.com/go-chi/chi"
)

// New return a new instance of kubac's router.
func New() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/ping", HandlePing)

	return r
}

// HandlePing is a handler returning 204 No Content
func HandlePing(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}
