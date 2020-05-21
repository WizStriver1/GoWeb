# install过程中无法成功

可以看到报错信息下有"cannot find package"提示，表示找不到依赖包，通过执行go get命令批量下载所有远程依赖包

```
go get -d -v gowebapp/...
```

报错信息：
```
G:\my\SomeTest\GoLang>go install gowebapp
src\gowebapp\vendor\app\shared\database\database.go:9:2: cannot find package "github.com/boltdb/bolt" in any of:
        G:\my\SomeTest\GoLang\src\gowebapp\vendor\github.com\boltdb\bolt (vendor tree)
        D:\Go\src\github.com\boltdb\bolt (from $GOROOT)
        G:\my\SomeTest\GoLang\src\github.com\boltdb\bolt (from $GOPATH)
        D:\Users\wizstriver\go\src\github.com\boltdb\bolt
src\gowebapp\vendor\app\shared\database\database.go:10:2: cannot find package "github.com/go-sql-driver/mysql" in any of:
        G:\my\SomeTest\GoLang\src\gowebapp\vendor\github.com\go-sql-driver\mysql (vendor tree)
        D:\Go\src\github.com\go-sql-driver\mysql (from $GOROOT)
        G:\my\SomeTest\GoLang\src\github.com\go-sql-driver\mysql (from $GOPATH)
        D:\Users\wizstriver\go\src\github.com\go-sql-driver\mysql
src\gowebapp\vendor\app\controller\notepad.go:12:2: cannot find package "github.com/gorilla/context" in any of:
        G:\my\SomeTest\GoLang\src\gowebapp\vendor\github.com\gorilla\context (vendor tree)
        D:\Go\src\github.com\gorilla\context (from $GOROOT)
        G:\my\SomeTest\GoLang\src\github.com\gorilla\context (from $GOPATH)
        D:\Users\wizstriver\go\src\github.com\gorilla\context
src\gowebapp\vendor\app\shared\session\session.go:6:2: cannot find package "github.com/gorilla/sessions" in any of:
        G:\my\SomeTest\GoLang\src\gowebapp\vendor\github.com\gorilla\sessions (vendor tree)
        D:\Go\src\github.com\gorilla\sessions (from $GOROOT)
        G:\my\SomeTest\GoLang\src\github.com\gorilla\sessions (from $GOPATH)
        D:\Users\wizstriver\go\src\github.com\gorilla\sessions
src\gowebapp\vendor\app\shared\recaptcha\recaptcha.go:7:2: cannot find package "github.com/haisum/recaptcha" in any of:
        G:\my\SomeTest\GoLang\src\gowebapp\vendor\github.com\haisum\recaptcha (vendor tree)
        D:\Go\src\github.com\haisum\recaptcha (from $GOROOT)
        G:\my\SomeTest\GoLang\src\github.com\haisum\recaptcha (from $GOPATH)
        D:\Users\wizstriver\go\src\github.com\haisum\recaptcha
src\gowebapp\vendor\app\shared\database\database.go:11:2: cannot find package "github.com/jmoiron/sqlx" in any of:
        G:\my\SomeTest\GoLang\src\gowebapp\vendor\github.com\jmoiron\sqlx (vendor tree)
        D:\Go\src\github.com\jmoiron\sqlx (from $GOROOT)
        G:\my\SomeTest\GoLang\src\github.com\jmoiron\sqlx (from $GOPATH)
        D:\Users\wizstriver\go\src\github.com\jmoiron\sqlx
src\gowebapp\vendor\app\controller\login.go:14:2: cannot find package "github.com/josephspurrier/csrfbanana" in any of:
        G:\my\SomeTest\GoLang\src\gowebapp\vendor\github.com\josephspurrier\csrfbanana (vendor tree)
        D:\Go\src\github.com\josephspurrier\csrfbanana (from $GOROOT)
        G:\my\SomeTest\GoLang\src\github.com\josephspurrier\csrfbanana (from $GOPATH)
        D:\Users\wizstriver\go\src\github.com\josephspurrier\csrfbanana
src\gowebapp\vendor\app\controller\notepad.go:14:2: cannot find package "github.com/julienschmidt/httprouter" in any of:
        G:\my\SomeTest\GoLang\src\gowebapp\vendor\github.com\julienschmidt\httprouter (vendor tree)
        D:\Go\src\github.com\julienschmidt\httprouter (from $GOROOT)
        G:\my\SomeTest\GoLang\src\github.com\julienschmidt\httprouter (from $GOPATH)
        D:\Users\wizstriver\go\src\github.com\julienschmidt\httprouter
src\gowebapp\vendor\app\route\route.go:16:2: cannot find package "github.com/justinas/alice" in any of:
        G:\my\SomeTest\GoLang\src\gowebapp\vendor\github.com\justinas\alice (vendor tree)
        D:\Go\src\github.com\justinas\alice (from $GOROOT)
        G:\my\SomeTest\GoLang\src\github.com\justinas\alice (from $GOPATH)
        D:\Users\wizstriver\go\src\github.com\justinas\alice
src\gowebapp\vendor\app\shared\passhash\passhash.go:4:2: cannot find package "golang.org/x/crypto/bcrypt" in any of:
        G:\my\SomeTest\GoLang\src\gowebapp\vendor\golang.org\x\crypto\bcrypt (vendor tree)
        D:\Go\src\golang.org\x\crypto\bcrypt (from $GOROOT)
        G:\my\SomeTest\GoLang\src\golang.org\x\crypto\bcrypt (from $GOPATH)
        D:\Users\wizstriver\go\src\golang.org\x\crypto\bcrypt
src\gowebapp\vendor\app\shared\database\database.go:12:2: cannot find package "gopkg.in/mgo.v2" in any of:
        G:\my\SomeTest\GoLang\src\gowebapp\vendor\gopkg.in\mgo.v2 (vendor tree)
        D:\Go\src\gopkg.in\mgo.v2 (from $GOROOT)
        G:\my\SomeTest\GoLang\src\gopkg.in\mgo.v2 (from $GOPATH)
        D:\Users\wizstriver\go\src\gopkg.in\mgo.v2
src\gowebapp\vendor\app\model\note.go:13:2: cannot find package "gopkg.in/mgo.v2/bson" in any of:
        G:\my\SomeTest\GoLang\src\gowebapp\vendor\gopkg.in\mgo.v2\bson (vendor tree)
        D:\Go\src\gopkg.in\mgo.v2\bson (from $GOROOT)
        G:\my\SomeTest\GoLang\src\gopkg.in\mgo.v2\bson (from $GOPATH)
        D:\Users\wizstriver\go\src\gopkg.in\mgo.v2\bson

G:\my\SomeTest\GoLang>
```