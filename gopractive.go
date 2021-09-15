package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

type IndexData struct {
	Title string
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析 url 傳遞的參數，對於 POST 則解析 HTTP 回應內容的主體（request body）
	//注意 : 如果沒有呼叫 ParseForm 方法，下面無法取得表單的資料
	fmt.Println(r.Form) //這些資訊是輸出到伺服器端的列印資訊
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	var name string
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
		fmt.Println(k)
		if k == "username" {
			name = strings.Join(v, "")
		}
	}
	fmt.Fprintf(w, "Hello "+name+" !") //這個寫入到 w 的是輸出到客戶端的
}

func test(w http.ResponseWriter, r *http.Request) {
	//tmpl := template.Must(template.ParseFiles("./index.gtpl"))
	//data := new(IndexData)
	//data.Title = "測試"
	//tmpl.Execute(w, data)
	fmt.Println("method:", r.Method) //取得請求的方法
	if r.Method == "GET" {
		t, _ := template.ParseFiles("index.gtpl")
		log.Println(t.Execute(w, nil))
	} else {
		//r.ParseForm()
		//請求的是登入資料，那麼執行登入的邏輯判斷
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}

func main() {
	http.HandleFunc("/", sayhelloName)
	http.HandleFunc("/index", test)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
