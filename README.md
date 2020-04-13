#　cloudgo

### 目录

```
➜  cloudgo git:(master) tree
.
├── conf
│   └── app.ini
├── main.go
├── middleware
├── models
│   ├── models.go
│   └── sql
│       └── blog.sql
├── pkg
│   ├── setting
│   │   └── setting.go
│   ├── status
│   │   ├── code.go
│   │   └── msg.go
│   └── util
│       └── pagination.go
├── routers
│   └── router.go
├── runtime
├── service
│   └── server.go
├── templates
└── test
    └── test.go

```

### 包依赖
```
github.com/go-ini/ini
github.com/Unknwon/com
github.com/gin-gonic/gin
github.com/jinzhu/gorm
// mysql驱动
github.com/go-sql-driver/mysql
// beego 表单　validation
go get -u github.com/astaxie/beego/validation
// jwt doc https://godoc.org/github.com/dgrijalva/jwt-go
go get -u github.com/dgrijalva/jwt-go
```