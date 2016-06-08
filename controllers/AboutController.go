package controllers

type AboutController struct {
	baseController
}

func (this *AboutController) Get(){
	this.TplName="About.html"
}

