package controllers

type AboutController struct {
	baseController
}

func (this *AboutController) Get(){
	this.Ctx.Redirect(302,"/resume")
	this.EnableRender = false
	//this.TplName="About.html"
}

