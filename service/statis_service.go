package service

import (
	"github.com/go-xorm/xorm"
	"irisdemo/model"
	"time"
)

type StatisService interface {
	GetAdminDailyCount(date string) int64
	GetUserDailyCount(date string) int64
	GetOrderDailyCount(date string) int64
}
func NewStatisService (db *xorm.Engine) StatisService{
	return &statisService{
		engine:db,
	}
}
type statisService struct {
	engine *xorm.Engine
}


func (this *statisService) GetAdminDailyCount(date string) int64{
	if date == "NaN-NaN-NaN" {
		date = time.Now().Format("2006-01-02")
	}
	startDate,err := time.Parse("2006-01-02",date)
	if err != nil{
		return 0
	}
	endDate := startDate.AddDate(0,0,1)
	result,err := this.engine.Where("create_time between ? and ? and status = 0",startDate.Format("2006-01-02 15:04:05"),endDate.Format("2006-01-02 15:04:05")).Count(model.Admin{})
	if err != nil{
		return 0
	}
	return result
}

func (this *statisService) GetUserDailyCount(date string) int64{
	if date == "NaN-NaN-NaN" {
		date = time.Now().Format("2006-01-02")
	}
	startDate,err := time.Parse("2006-01-02",date)
	if err != nil{
		return 0
	}
	endDate := startDate.AddDate(0,0,1)
	result,err := this.engine.Where("register_time between ? and ? and del_flag = 0",startDate.Format("2006-01-02 15:04:05"),endDate.Format("2006-01-02 15:04:05")).Count(model.User{})
	if err != nil{
		return 0
	}
	return result
}

func (this *statisService) GetOrderDailyCount(date string) int64{
	if date == "NaN-NaN-NaN" {
		date = time.Now().Format("2006-01-02")
	}
	startDate,err := time.Parse("2006-01-02",date)
	if err != nil{
		return 0
	}
	endDate := startDate.AddDate(0,0,1)
	result,err := this.engine.Where("time between ? and ? and del_flag = 0",startDate.Format("2006-01-02 15:04:05"),endDate.Format("2006-01-02 15:04:05")).Count(model.Order{})
	if err != nil{
		return 0
	}
	return result
}