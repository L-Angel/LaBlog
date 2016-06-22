package controllers


type AdminAddPictureController struct {
	baseController
}

func (this *AdminAddPictureController) Get(){
	this.TplName="admin/AdminAddPicture.html"
}

