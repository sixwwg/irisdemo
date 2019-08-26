package controller

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
	"irisdemo/service"
	"irisdemo/utils"
	"ulog_service/iris/mvc"
)

type UserController struct {
	Context iris.Context
	Service service.UserService
	Session *sessions.Session
}
/**
 * 获取用户总数
 * user/count
 */
func (this *UserController) GetCount() mvc.Result{
	count,err := this.Service.GetUserCount()
	if err != nil{
		return mvc.Response{
			Object:map[string]interface{}{
				"status"  : utils.RESPONSE_FAIL,
				"message" : utils.Recode2Text(utils.RESMSG_ERROR_USERLIST),
			},
		}
	}
	return mvc.Response{
		Object:map[string]interface{}{
			"status"  : utils.RESPONSE_OK,
			"count"   : count,
		},
	}
}
