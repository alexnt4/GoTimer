package utils

import (
	"fmt"
	"html/template"
	"net/http"
)

func ParseTemplateFiles(w http.ResponseWriter, templateName string, context any, funcToTemplate template.FuncMap) {
	ts, er := template.ParseFiles(files...)
	if er != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		fmt.Println("Error parsing template files", er)
		return
	}

	if funcToTemplate != nil {
		ts = ts.Funcs(funcToTemplate)
	}

	error := ts.ExecuteTemplate(w, templateName, context)
	if error != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		fmt.Println("Error executing template", error)
	}
}
