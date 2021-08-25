package main

import (
	"fmt"
	"github.com/13808796047/go-gin-example/models"
	_ "github.com/13808796047/go-gin-example/models"
	"github.com/13808796047/go-gin-example/pkg/gredis"
	"github.com/13808796047/go-gin-example/pkg/logging"
	"github.com/13808796047/go-gin-example/pkg/setting"
	"github.com/13808796047/go-gin-example/routers"
	"net/http"
)

func main() {
	setting.Setup()
	models.Setup()
	logging.Setup()
	gredis.Setup()
	router := routers.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HttpPort),
		Handler:        router,
		ReadTimeout:    setting.ServerSetting.ReadTimeout,
		WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
