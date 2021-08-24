package api

import (
	"github.com/13808796047/go-gin-example/models"
	"github.com/13808796047/go-gin-example/pkg/e"
	"github.com/13808796047/go-gin-example/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type auth struct {
	Username string `valid:"Required;MaxSize(50)"`
	Password string `valid:"Required;MaxSize(50)"`
}

func GetAuth(c *gin.Context)  {
	username := c.Query("username")
	password := c.Query("password")
	a := auth{Username: username,Password: password}
	valid := validation.Validation{}
	ok,_ := valid.Valid(&a)
	data := make(map[string]interface{})
	code := e.INVALID_PARAMS
	if ok {
		isExist := models.CheckAuth(username,password)
		if isExist{
			token,err := util.GenerateToken(username,password)
			if err !=nil {
				code = e.ERROR_AUTH
			}else{
				data["token"] = token
				code = e.SUCCESS
			}
		}else {
			code = e.ERROR_AUTH
		}
	}else{
		for _,err := range valid.Errors{
			log.Println(err.Key,err.Message)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : data,
	})
}