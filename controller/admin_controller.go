package controller

import (
	"encoding/json"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"irisdemo/model"
	"irisdemo/service"
	"irisdemo/utils"
)

type AdminController struct {
	Context iris.Context
	Service service.AdminService
	Session *sessions.Session
}

const (
	ADMIN_TABLE_NAME 	=  "admin"
	ADMIN 				=  "admin"
)

type AdminLogin struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

/**
 * 管理员退出功能
 * 请求类型 Get
 * 请求url : admin/logout
 */
func (this *AdminController) GetLogout() mvc.Result{
	//删除session
	this.Session.Delete(ADMIN)
	return mvc.Response{
		Object:map[string]interface{}{
			"status"  : utils.RESPONSE_OK,
			"success" : utils.Recode2Text(utils.RESMSG_SUCCESSSIGNOUT),
		},
	}
}


/**
 * 管理员登录功能
 * 请求类型 Post
 * 请求url : admin/login
 */
func (this *AdminController) PostLogin(ctx iris.Context) mvc.Result{
	username := ctx.URLParamDefault("username","")
	password := ctx.URLParamDefault("password","")
	//var adminLogin AdminLogin
	//err := this.Context.ReadJSON(&adminLogin)
	//if err != nil{
	//	fmt.Println("解析出错",err)
	//}
	//数据参数校验
	if username == "" || password == ""{
		return mvc.Response{
			Object:map[string]interface{}{
				"status"  : 0,
				"error"   : "登录失败",
				"message" : "登录名或者密码为空，输入尼玛呢",
			},
		}
	}
	//查询数据库
	admin,exist := this.Service.GetAdminByNameAndPassword(username,password)
	if !exist{
		return mvc.Response{
			Object:map[string]interface{}{
				"status"  : 0,
				"error"   : "登录失败",
				"message" : "用户名或密码错误，请重新输入",
			},
		}
	}
	adminByte,_ := json.Marshal(admin)
	this.Session.Set(ADMIN,adminByte)
	return mvc.Response{
		Object:map[string]interface{}{
			"status"  : 1,
			"success" : "登录成功",
			"message" : "管理员登录成功",
		},
	}
}

/**
 * 获取管理员信息
 * 请求类型 Get
 * 请求url : admin/info
 */
func (this *AdminController) GetInfo() mvc.Result{
	//从session中获取用户信息
	userByte := this.Session.Get(ADMIN)

	//session为空
	if userByte == nil{
		return mvc.Response{
			Object:map[string]interface{}{
				"status" : utils.RESPONSE_NOLOGIN,
				"type"	  : utils.RESMSG_UNLOGIN,
				"message" : utils.Recode2Text(utils.RESMSG_UNLOGIN),
			},
		}
	}
	var admin model.Admin
	err := json.Unmarshal(userByte.([]byte),&admin)
	//session信息解析失败
	if err != nil{
		return mvc.Response{
			Object:map[string]interface{}{
				"status" : utils.RESPONSE_NOLOGIN,
				"type"	  : utils.RESMSG_UNLOGIN,
				"message" : utils.Recode2Text(utils.RESMSG_UNLOGIN),
			},
		}
	}
	//session解析成功
	return mvc.Response{
		Object:map[string]interface{}{
			"status" : utils.RESPONSE_OK,
			"data"   : admin.AdminToRespDesc(), //信息
		},
	}
}

/**
 * 获取管理员数量
 * Get
 * url /admin/count
 */

func (this *AdminController) GetCount() mvc.Result{
	count,err := this.Service.GetAdminCount()
	if err != nil{
		return mvc.Response{
			Object:map[string]interface{}{
				"status" : utils.RESPONSE_FAIL,
				"message" : utils.Recode2Text(utils.RESMSG_ERRORADMINCOUNT),
				"count"   : 0,
			},
		}
	}
	return mvc.Response{
		Object:map[string]interface{}{
			"status" : utils.RESPONSE_OK,
			"count"   : count,
		},
	}
}