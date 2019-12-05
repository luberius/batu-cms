package controllers

import (
	"net/http"
	"batu-cms/core/view"
	"batu-cms/domain/user"
)

// Index controller
func Index(w http.ResponseWriter, r *http.Request)  {
	user.Hello()
	view.Render(w, "World", "views/home.tmpl", "views/layout/header.tmpl", "views/layout/footer.tmpl")
}