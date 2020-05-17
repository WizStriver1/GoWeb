package main

import (
	"fmt"
    "html/template"
	"net/http"
	"strings"
	"log"
	"mymath"
	"strconv"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	form := r.Form
	path := r.URL.Path
	scheme := r.URL.Scheme
	urlLong := r.Form["url_long"]
	method := r.Method
	fmt.Println("Formdata:")
	fmt.Println(form)
	fmt.Println("path", path)
	fmt.Println("scheme", scheme)
	fmt.Println("url_long", urlLong)
    fmt.Println("method:", method) // 获取请求的方法
	var value float64
	for k, v := range r.Form {
		fmt.Println("key", k)
		fmt.Println("val:", strings.Join(v, ""))
		if k == "test" {
			val, err := strconv.ParseFloat(strings.Join(v, ""), 64)
			fmt.Println()
			if err == nil {
				value = mymath.Sqrt(val)
			}
		}
	}
	fmt.Println("===========================")
	fmt.Fprintf(w, "hello astaxie! The value is %v", value)
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		log.Println(t.Execute(w, nil))
	} else {
		err := r.ParseForm()
		if err != nil {
			log.Fatal("ParseForm", err)
		}
		fmt.Println("username: ", r.Form["username"])
		fmt.Println("password: ", r.Form["password"])
		t2, _ := template.ParseFiles("success.gtpl")
		log.Println(t2.Execute(w, nil))
	}
}

func main() {
	fmt.Println("Please visit http://localhost:9090")
	http.HandleFunc("/", sayHelloName)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
