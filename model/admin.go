package model

import "time"

type Admin struct {
	AdminId int64 `xorm:"pk autoincr" json:"id"`
	AdminName string `xorm:"varchar(32)" json:"admin_name"`
	CreateTime time.Time `xorm:"DateTime" json:"create_time"`
	Status int `xorm:"default 0" json:"status"`
	Avatar string `xorm:"varchar(255)" json:"avatar"`
	Pwd string `xorm:"varchar(255)" json:"pwd"`
	CityName string `xorm:"varchar(12)" json:"city_name"`
	CityId int64 `xorm:"index" json:"city_id"`
}

func (this *Admin) AdminToRespDesc() interface{}{
	return map[string]interface{}{
		"user_name" : this.AdminName,
		"id"		: this.AdminId,
		"create_time" : this.CreateTime,
		"status"	: this.Status,
		"avatar" 	: this.Avatar,
		"city"		: this.CityName,
		"admin"		: "管理员",
	}
}