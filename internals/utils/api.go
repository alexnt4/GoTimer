package utils

import (
	"fmt"
	"net/http"
	"text/template"
)

func CheckIfPath(w http.ResponseWriter, r *http.Request, path string) {
	if r.URL.Path != path {
		fmt.Println("Error: "+r.URL.Path, path)
		http.NotFound(w, r)
		return
	}
}

func IsValidHTTPMehod(method string, acceptedMethod string, w http.ResponseWriter) bool {
	if method == acceptedMethod {
		w.Header().Set("Allow", acceptedMethod)
		return true
	}

	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	return false
}

var (
	EmptyFuncMap = template.FuncMap{}
	EmptyStruct  = struct{}{}
)
