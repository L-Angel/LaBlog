package controllers


type CategoryController struct {
	baseController
}

func (this *CategoryController) Get(){
	this.TplName="Category.html"
}