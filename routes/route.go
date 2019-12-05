package routes

import (
	"github.com/go-chi/chi"
	"batu-cms/controllers"
)

// InitRoute function for initiate routes
func InitRoute(r *chi.Mux) {
	r.Get("/", controllers.Index)
}