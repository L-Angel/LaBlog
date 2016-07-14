package models

import (
//"fmt"
//"strings"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
_	 "github.com/go-sql-driver/mysql"
)

func RegisterDataBase() {
	var DbUser = beego.AppConfig.String("dbusername")
	var DbPass = beego.AppConfig.String("dbpassword")
	var DbHosts = beego.AppConfig.String("dbhosts")
	//var DbType = beego.AppConfig.String("dbtype")
	var DbName = beego.AppConfig.String("dbname")
        var DbPort=beego.AppConfig.String("dbport")
	var ConnString = DbUser + ":" + DbPass + "@tcp(" + DbHosts +":"+DbPort+ ")/" + DbName
	beego.Error(ConnString)
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", ConnString)
	orm.Debug=true
}
