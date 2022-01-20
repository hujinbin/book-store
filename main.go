package main

import (
	"net/http"
	"text/template"
	"controller"
)

func Indexhandler(w http.ResponseWriter, r *http.Request)  {
	t := template.Must(template.ParseFiles("page/index.html"))
	t.Execute(w, "")
}

func main()  {
	http.HandleFunc("/", Indexhandler)
	http.Handle("/files/",http.StripPrefix("/files/",http.FileServer(http.Dir("files"))))
	http.Handle("/page/",http.StripPrefix("/page/",http.FileServer(http.Dir("page"))))
	http.HandleFunc("/login", controller.Login)
	http.ListenAndServe(":6060", nil)
} 