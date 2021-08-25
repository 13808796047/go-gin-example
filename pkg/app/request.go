package app

import (
	"github.com/13808796047/go-gin-example/pkg/logging"
	"github.com/astaxie/beego/validation"
)

// 返回错误
func MarkErrors(errors []*validation.Error) {
	for _, err := range errors {
		logging.Info(err.Key, err.Message)
	}
	return
}
