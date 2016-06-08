package controllers

type ErrorController struct {
	baseController
}

func (this *ErrorController) Error404(){
	this.TplName="error/404.html"
}
func (this *ErrorController) Error500(){
	this.TplName="error/500.html"
}
func (this *ErrorController) ErrorDb(){
	this.TplName="error/ErrorDb.html"
}