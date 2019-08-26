package service

import (
	"github.com/go-xorm/xorm"
	"irisdemo/model"
)

type AdminService interface {
	//通过登录用户名和密码查询数据库，核对登录是否合法
	GetAdminByNameAndPassword(username,password string) (model.Admin,bool)
	//获取管理员总数
	GetAdminCount()(int64,error)
}

func NewAdminService(db *xorm.Engine) AdminService{
	return &adminService{
		engine:db,
	}
}

type adminService struct {
	engine *xorm.Engine
}

func (this *adminService) GetAdminByNameAndPassword(username, password string) (model.Admin, bool) {
	var admin model.Admin
	if _,err := this.engine.Where("admin_name = ? and pwd = ?",username,password).Get(&admin);err != nil{
		panic(err)
		return admin,false
	}
	return admin,admin.AdminId != 0
}

func (this *adminService) GetAdminCount() (int64, error) {
	count,err := this.engine.Count(new(model.Admin))
	if err != nil{
		panic(err)
		return 0,err
	}
	return count,nil
}

