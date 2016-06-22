package controllers

import (

)

type AdminAddUserController struct {
	baseController
}

func (this *AdminAddUserController) Get(){
	this.TplName="admin/AdminAddUser.html"
}

