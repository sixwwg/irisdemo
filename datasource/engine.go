package datasource

import (
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"irisdemo/model"
	"xorm.io/core"
)

//todo 数据库引擎

func NewMysqlEngine() (*xorm.Engine){
	engine,err :=xorm.NewEngine("mysql","root:123456@/irisdemo?charset=utf8")
	if err != nil{
		err = errors.New("数据库连接失败 " + err.Error())
		panic(err)
		return nil
	}
	//设置表字段名为驼峰式命名映射规则
	engine.SetMapper(core.SnakeMapper{})
	//根据结构体同步表结构
	err = engine.Sync2(
		new(model.User),
		new(model.MonitorParam),
		new(model.Admin),
	)
	if err != nil{
		err = errors.New("数据库同步失败 " + err.Error())
		panic(err)
		return nil
	}
	//设置显示sql语句
	engine.ShowSQL(true)
	engine.SetMaxOpenConns(10)
	return engine
}