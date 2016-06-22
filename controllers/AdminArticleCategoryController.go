package controllers

import (

)

type AdminArticleCategoryController struct {
	baseController
}

func (this *AdminArticleCategoryController) Get(){
	this.TplName="admin/AdminArticleCategory.html"
}
