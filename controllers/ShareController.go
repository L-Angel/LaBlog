package controllers

import (

)

type ShareController struct {
	baseController
}

func (this *ShareController) Get(){
	this.TplName="Share.html"
}