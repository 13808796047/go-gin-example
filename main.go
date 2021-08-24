package main

import (
	"fmt"
	"github.com/13808796047/go-gin-example/pkg/setting"
	"github.com/13808796047/go-gin-example/routers"
	"net/http"
	_ "github.com/13808796047/go-gin-example/models"
)

func main() {
	router := routers.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
