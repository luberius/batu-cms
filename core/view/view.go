package view

import (
	"fmt"
	"net/http"
	"text/template"
)

//Render function for rendering view tamplate
func Render(w http.ResponseWriter, data interface{}, tmpl ...string) {
	tmpls := append([]string{"core/template/root.tmpl"}, tmpl...)
	t, err := template.ParseFiles(tmpls...)

	if err != nil {
		fmt.Fprint(w, "Error Parsing template: "+err.Error())
	} else {
		err = t.ExecuteTemplate(w, "root", data)

		if err != nil {
			fmt.Fprint(w, "Error Execute template: "+err.Error())
		}
	}
}
