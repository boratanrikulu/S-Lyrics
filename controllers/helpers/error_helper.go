package helpers

import (
	"log"
	"net/http"
	"html/template"
)

type PageData struct {
	ErrorMessage string
}

func ErrorPage(error string, w http.ResponseWriter) {
	p := PageData {
		ErrorMessage: error,
	}
	files := GetTemplateFiles("./views/error.html")
	tpml := template.Must(template.ParseFiles(files...))
	err := tpml.Execute(w, p)
	if err != nil {
		log.Fatal("Error occur while rendering error page: %v", error)
	}
}