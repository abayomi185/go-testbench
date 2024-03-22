package api

import (
	"net/http"

	"github.com/abayomi185/go-testbench/cmd/hello-world"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func api() {

	helloworld.Hello()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {})
	http.ListenAndServe(":3000", r)
}
