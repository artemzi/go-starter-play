package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // nolint
	fmt.Println(r.Form)
	fmt.Println("Path: ", r.URL.Path)
	fmt.Println("Scheme: ", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("Key: ", k)
		fmt.Println("Value: ", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello user!") // write data into response
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method: ", r.Method)
	if r.Method == "GET" {
		curTime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(curTime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("tpl/login.gtpl")
		t.Execute(w, token)
	} else {
		r.ParseForm()
		token := r.Form.Get("token")
		if token != "" {
			// проверяем валидность токена
		} else {
			// если нет токена, возвращаем ошибку
		}
		fmt.Println("username length:", len(r.Form["username"][0]))
		fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username"))) // печатаем на стороне сервера
		fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))
		template.HTMLEscape(w, []byte(r.Form.Get("username"))) // отвечаем клиенту
	}
}

func main() {
	http.HandleFunc("/", sayHelloName)
	http.HandleFunc("/login", login)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
