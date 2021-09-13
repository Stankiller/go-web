package main

import (
	"html/template"
	"log"
	"net/http"
)

type IndexData struct {
	Title    string
	Content  string
	Account  string
	Password string
}

func test(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./index.html"))
	data := new(IndexData)
	data.Title = "測試"
	data.Content = "測試"
	data.Account = "帳號"
	data.Password = "密碼"
	tmpl.Execute(w, data)
}

func main() {
	http.HandleFunc("/", test)
	http.HandleFunc("/index", test)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
