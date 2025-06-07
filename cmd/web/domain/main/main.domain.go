package mainDomain

import (
	"GoTimer/internals/models"
	"GoTimer/internals/utils"
	"fmt"
	"net/http"
	"text/template"
)

func MainView(w http.ResponseWriter, r *http.Request) {
	utils.CheckIfPath(w, r, models.RoutesIntance.MAIN)

	files := []string{
		"ui/html/base.html",
		"ui/html/pages/main.tmpl.html",
	}

	ts, er := template.ParseFiles(files...)
	if er != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		fmt.Println("Error parsing template files", er)
	}

	err := ts.ExecuteTemplate(w, "base", utils.EmptyStruct)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		fmt.Println("Error executing template")
	}
}
