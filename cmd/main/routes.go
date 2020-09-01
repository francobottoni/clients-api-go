package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func Routes() *chi.Mux {
	mux := chi.NewMux()

	//globals middlewares
	mux.Use(
		middleware.Logger,
		middleware.Recoverer,
	)

	mux.Post("/clients", nil)

	mux.Get("/hello", hellowHandler)

	return mux
}

func hellowHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content - Type", "application/json")
	w.Header().Set("done-by", "franco")

	res := map[string]interface{}{"message": "hello world"}

	_ = json.NewEncoder(w).Encode(res)
}
