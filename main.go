package main

import (
	"github.com/HashCell/golang/cloudgo/routers"
	"net/http"
	"fmt"
	"github.com/HashCell/golang/cloudgo/pkg/setting"
)

func main() {
	router := routers.InitRouter()
	server := &http.Server{
		Addr:fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:router,
		ReadTimeout:setting.ReadTimeout,
		WriteTimeout:setting.WriteTimeout,
		MaxHeaderBytes:1<<20,
	}

	server.ListenAndServe()
}
