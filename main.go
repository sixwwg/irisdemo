package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	config2 "irisdemo/config"
	"irisdemo/controller"
	"irisdemo/datasource"
	"irisdemo/service"
	"time"
)

/**
 * 程序主入口
 */
func main() {
	app := newApp()
	configuration(app)
	mvcHandle(app)
	config := config2.InitConfig()
	addr := ":" + config.Port
	app.Run(
		iris.Addr(addr),                               //在端口9000进行监听
		iris.WithoutServerError(iris.ErrServerClosed), //无服务错误提示
		iris.WithOptimizations,                        //对json数据序列化更快的配置
	)


}

//构建application
func newApp() *iris.Application{
	app := iris.New()
	//设置日志级别
	app.Logger().SetLevel("debug")
	//注册静态资源
	app.StaticWeb("/static","./static")
	app.StaticWeb("/manage/static","./static")
	//注册视图文件
	app.RegisterView(iris.HTML("./static",".html"))
	app.Get("/", func(context context.Context) {
		context.View("index.html")
	})
	return app
}

/**
 * 项目设置
 */
func configuration(app *iris.Application){
	//配置 字符编码
	app.Configure(iris.WithConfiguration(iris.Configuration{
		Charset:"UTF-8",
	}))
	//错误配置
	app.OnErrorCode(iris.StatusNotFound, func(context context.Context) {
		_,_ = context.JSON(iris.Map{
			"errmsg" 	: iris.StatusNotFound,
			"msg"		: "not found",
			"data"		: iris.Map{},
		})
	})
	app.OnErrorCode(iris.StatusInternalServerError, func(context context.Context) {
		_,_ =  context.JSON(iris.Map{
			"errmsg" 	: iris.StatusInternalServerError,
			"msg"		: "internal error",
			"data"		: iris.Map{},
		})
	})
}

/**
 * 路由设置
 */
func mvcHandle(app *iris.Application)  {
	//启用session
	sessionManager := sessions.New(sessions.Config{
		Cookie:"sessionCookie",
		Expires:24*time.Hour,
	})
	engine := datasource.NewMysqlEngine()

	//管理员模块功能
	adminService := service.NewAdminService(engine)
	//设置一个路由
	admin := mvc.New(app.Party("/admin"))
	admin.Register(
		adminService,
		sessionManager.Start,
	)
	//为上述的admin路由的请求设置一个处理器，处理器是自定义的
	admin.Handle(new(controller.AdminController))

	userService := service.NewUserService(engine)
	user := mvc.New(app.Party("/v1/user"))
	user.Register(
		userService,
		sessionManager.Start,
	)
	user.Handle(new(controller.UserController))


}
