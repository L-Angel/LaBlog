package controllers

import (

)

type AdminIndexController struct  {
	baseController
}


func (this *AdminIndexController) Get(){
	this.TplName="admin/AdminIndex.html"
}

