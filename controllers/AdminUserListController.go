package controllers

import (

)

type  AdminUserListController  struct{
	baseController
}

func (this *AdminUserListController) Get(){
	this.TplName="admin/AdminUserList.html"
}