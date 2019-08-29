package controller

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
	"irisdemo/service"
)

type StatisController struct {
	Context iris.Context
	Service service.StatisService
	Session sessions.Session
}

//func