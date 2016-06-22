package controllers

type RegisterController struct {
	baseController
}

func (this *RegisterController) Get(){
	this.TplName="admin/Register.html"
}
