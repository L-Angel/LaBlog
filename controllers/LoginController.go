package controllers

type LoginController struct  {
	baseController
}

func (this *LoginController) Get(){
	this.TplName="Login.html"
}
