package controllers

import (
)

type LinksController struct {
	baseController
}

func (this *LinksController) Get() {
	this.TplName="Links.html"
}


