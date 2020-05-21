# http://localhost:9090/login

```

G:\my\SomeTest\GoLang\GoWeb\bin>webServer.exe
Please visit http://localhost:9090
method: GET
2020/05/21 13:38:10 http: panic serving [::1]:58759: runtime error: invalid memory address or nil pointer dereference
goroutine 6 [running]:
net/http.(*conn).serve.func1(0xc00004f040)
        D:/Go/src/net/http/server.go:1772 +0x140
panic(0x76fac0, 0xab46e0)
        D:/Go/src/runtime/panic.go:975 +0x3f1
html/template.(*Template).escape(0x0, 0x0, 0x0)
        D:/Go/src/html/template/template.go:95 +0x42
html/template.(*Template).Execute(0x0, 0x855620, 0xc00021a000, 0x0, 0x0, 0x855760, 0xc00008e000)
        D:/Go/src/html/template/template.go:119 +0x36
main.login(0x85af60, 0xc00021a000, 0xc000136000)
        D:/Users/wizstriver/go/src/webServer/main.go:46 +0x3e3
net/http.HandlerFunc.ServeHTTP(0x7eeed0, 0x85af60, 0xc00021a000, 0xc000136000)
        D:/Go/src/net/http/server.go:2012 +0x4b
net/http.(*ServeMux).ServeHTTP(0xac4320, 0x85af60, 0xc00021a000, 0xc000136000)
        D:/Go/src/net/http/server.go:2387 +0x1ac
net/http.serverHandler.ServeHTTP(0xc00012c000, 0x85af60, 0xc00021a000, 0xc000136000)
        D:/Go/src/net/http/server.go:2807 +0xaa
net/http.(*conn).serve(0xc00004f040, 0x85b5e0, 0xc00003a340)
        D:/Go/src/net/http/server.go:1895 +0x873
created by net/http.(*Server).Serve
        D:/Go/src/net/http/server.go:2933 +0x363
method: GET
2020/05/21 13:38:10 http: panic serving [::1]:58760: runtime error: invalid memory address or nil pointer dereference
goroutine 18 [running]:
net/http.(*conn).serve.func1(0xc000188000)
        D:/Go/src/net/http/server.go:1772 +0x140
panic(0x76fac0, 0xab46e0)
        D:/Go/src/runtime/panic.go:975 +0x3f1
html/template.(*Template).escape(0x0, 0x0, 0x0)
        D:/Go/src/html/template/template.go:95 +0x42
html/template.(*Template).Execute(0x0, 0x855620, 0xc00021a0e0, 0x0, 0x0, 0x855760, 0xc000206150)
        D:/Go/src/html/template/template.go:119 +0x36
main.login(0x85af60, 0xc00021a0e0, 0xc00019a000)
        D:/Users/wizstriver/go/src/webServer/main.go:46 +0x3e3
net/http.HandlerFunc.ServeHTTP(0x7eeed0, 0x85af60, 0xc00021a0e0, 0xc00019a000)
        D:/Go/src/net/http/server.go:2012 +0x4b
net/http.(*ServeMux).ServeHTTP(0xac4320, 0x85af60, 0xc00021a0e0, 0xc00019a000)
        D:/Go/src/net/http/server.go:2387 +0x1ac
net/http.serverHandler.ServeHTTP(0xc00012c000, 0x85af60, 0xc00021a0e0, 0xc00019a000)
        D:/Go/src/net/http/server.go:2807 +0xaa
net/http.(*conn).serve(0xc000188000, 0x85b5e0, 0xc000192000)
        D:/Go/src/net/http/server.go:1895 +0x873
created by net/http.(*Server).Serve
        D:/Go/src/net/http/server.go:2933 +0x363
method: GET
2020/05/21 13:38:10 http: panic serving [::1]:58761: runtime error: invalid memory address or nil pointer dereference
goroutine 36 [running]:
net/http.(*conn).serve.func1(0xc00024e000)
        D:/Go/src/net/http/server.go:1772 +0x140
panic(0x76fac0, 0xab46e0)
        D:/Go/src/runtime/panic.go:975 +0x3f1
html/template.(*Template).escape(0x0, 0x0, 0x0)
        D:/Go/src/html/template/template.go:95 +0x42
html/template.(*Template).Execute(0x0, 0x855620, 0xc0000c8000, 0x0, 0x0, 0x855760, 0xc00008e120)
        D:/Go/src/html/template/template.go:119 +0x36
main.login(0x85af60, 0xc0000c8000, 0xc0000ba000)
        D:/Users/wizstriver/go/src/webServer/main.go:46 +0x3e3
net/http.HandlerFunc.ServeHTTP(0x7eeed0, 0x85af60, 0xc0000c8000, 0xc0000ba000)
        D:/Go/src/net/http/server.go:2012 +0x4b
net/http.(*ServeMux).ServeHTTP(0xac4320, 0x85af60, 0xc0000c8000, 0xc0000ba000)
        D:/Go/src/net/http/server.go:2387 +0x1ac
net/http.serverHandler.ServeHTTP(0xc00012c000, 0x85af60, 0xc0000c8000, 0xc0000ba000)
        D:/Go/src/net/http/server.go:2807 +0xaa
net/http.(*conn).serve(0xc00024e000, 0x85b5e0, 0xc0000b8000)
        D:/Go/src/net/http/server.go:1895 +0x873
created by net/http.(*Server).Serve
        D:/Go/src/net/http/server.go:2933 +0x363

```