package model

import "time"

type MonitorParam struct {
	Id int `xorm:"pk autoincr"`  //主键自增 id
	Name string `xorm:"varchar(20)"`
	TaskId int `xorm:"int notnull"`
	TaskType int `xorm:"tinyint(2)"`     //枚举类型
	AlertType int `xorm:"tinyint(2)"`    //枚举类型
	Dbrps string `xorm:"varchar(100)"`
	MeasurementsTags string `xorm:"varchar(100)"`
	Fields string `xorm:"varchar(20)"`
	Function int `xorm:"tinyint(3)"`      //枚举类型
	Where string `xorm:"varchar(20)"`
	Condition string `xorm:"varchar(4)"`
	Threshold int `xorm:"int"`
	Period string `xorm:"varchar(20)"`
	Every string `xorm:"varchar(20)"`
	Info string `xorm:"varchar(100)"`
	Message string `xorm:"varchar(100)"`
	PostUrl string `xorm:"varchar(50)"`
	CreatedTime time.Time `xorm:"created"`
	UpdatedTime time.Time `xorm:"updated"`
}