package service

import (
	"errors"
	"github.com/alecthomas/log4go"
	"github.com/go-xorm/xorm"
	"irisdemo/model"
)

type UserService interface {
	GetUserCount() (int64,error)
	GetUserList(offset, limit int) []*model.User
}
func NewUserService (db *xorm.Engine) UserService{
	return &userService{
		engine:db,
	}
}

type userService struct {
	engine *xorm.Engine
}

func (this *userService) GetUserCount() (int64,error){
	count,err := this.engine.Count(new(model.User))
	if err != nil{
		err = errors.New("获取用户总数出错 " + err.Error())
		log4go.Error(err)
		return 0,err
	}
	return count,nil
}

func (this *userService) GetUserList(offset,limit int) []*model.User{
	var userList []*model.User
	err :=this.engine.Where("del_flag = ?",0).Limit(limit,offset).Find(&userList)
	if err != nil{
		err = errors.New("获取用户列表出错 " + err.Error())
		log4go.Error(err)
		return nil
	}
	return userList
}
