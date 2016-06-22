package controllers

import (

)

type  AdminArticleListController  struct{
	baseController
}

func (this *AdminArticleListController) Get(){
	this.TplName="admin/AdminArticleList.html"
}
