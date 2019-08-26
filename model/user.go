package model

type User struct {
	Id int `xorm:"pk autoincr"`  //主键自增 id
	Name string `xorm:"varchar(24)"`
	PhoneNum string `xorm:"varchar(20)"`
	Email int `xorm:"varchar(20) notnull"`
	City string `xorm:"-"` //不映射该字段
}

