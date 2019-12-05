package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"batu-cms/domain/user"
)

// Index controller
func Index(w http.ResponseWriter, r *http.Request)  {
	user.Hello()
	t, err := template.ParseFiles("views/home.tmpl", "views/layout/header.tmpl", "views/layout/footer.tmpl")

	if err != nil {
		fmt.Print(err.Error())
	} else {
		t.ExecuteTemplate(w, "layout", "")
	}
}