package main

import (
	"fmt"
	"html/template"
	"log"
	"mymath"
	"net/http"
	"strconv"
	"strings"
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
	fmt.Println("method:", r.Method) // 获取请求的方法
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		log.Println(t.Execute(w, nil))
	} else {
		err := r.ParseForm() // 解析 url 传递的参数，对于 POST 则解析响应包的主体（request body）
		if err != nil {
			// handle error http.Error() for example
			log.Fatal("ParseForm: ", err)
		}
		// 请求的是登录数据，那么执行登录的逻辑判断
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
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
