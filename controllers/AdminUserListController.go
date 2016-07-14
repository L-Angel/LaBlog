package controllers

import (
	"github.com/astaxie/beego"
	"github.com/L-Angel/LaBlog/models"
)

type  AdminUserListController  struct{
	baseController
}

func (this *AdminUserListController) Get(){
	this.Data["xsrf_token"]=this.XSRFToken()
	query:= map[string]string{}
	fields := make([]string,0)
	sortby:=[]string{"Id"}
	order := []string{"asc"}
	Users,err := models.GetAllUsers(query,fields , sortby,order,0,10000)
	//beego.Info(Blogs)
	if err !=nil{
		beego.Error(err)
	}
	this.Data["Users"]=Users
	this.TplName="admin/AdminUserList.html"
}