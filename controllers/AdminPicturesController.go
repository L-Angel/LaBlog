package controllers

type AdminPicturesController struct {
	baseController
}

func (this *AdminPicturesController) Get(){
	this.TplName="admin/AdminPictures.html"
}

