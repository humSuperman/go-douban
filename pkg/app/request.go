package app

import (
	"admin/pkg/logging"
	"github.com/astaxie/beego/validation"
)

func MarkErrors(errors []*validation.Error){
	for _,err := range errors{
		logging.Error(err.Key,err.Message)
	}
	return
}
