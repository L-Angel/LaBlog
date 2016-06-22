package controllers

import (

)

type ResumeController struct{
	baseController
}

func (this *ResumeController) Get(){
	this.TplName="Resume.html"
}
