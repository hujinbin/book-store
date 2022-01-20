package main

import (
	"net/http"
	"text/template"
)

func Indexhandler(w http.ResponseWriter, r*http.Request)  {
	t:= template.Must(template.ParseFiles("index.html"))
	http.Handle("/files/",http.StripPrefix("/files/",http.FileServer(http.Dir("/files"))))
	http.Handle("/page/",http.StripPrefix("/page/",http.FileServer(http.Dir("/page"))))
	t.Execute(w, "")
}

func main()  {
	http.HandleFunc("/main",Indexhandler)

	http.ListenAndServe(":6060", nil)
} 