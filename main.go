package main

import (
	"net/http"
	"github.com/go-chi/chi"
	"batu-cms/routes"
)

func main() {
	chiRouter := chi.NewRouter()
	routes.InitRoute(chiRouter)
	http.ListenAndServe(":3000", chiRouter)
}