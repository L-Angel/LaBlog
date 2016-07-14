package controllers

import (
	"github.com/L-Angel/LaBlog/models"
	"github.com/L-Angel/LaBlog/tools"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/utils/captcha"
)

func init(){
	store := cache.NewMemoryCache()
	cpt=captcha.NewWithFilter("/captcha/",store)
}

type RegisterController struct {
	baseController
}

func (this *RegisterController) Get(){
	this.Data["xsrf_token"]=this.XSRFToken()
	this.TplName="admin/Register.html"
}

func (this *RegisterController) Post(){
	username:=this.GetString("username")
	password:=tools.AES(this.GetString("password"))
	email:=this.GetString("email")
	tel:=this.GetString("tel")
	address:=this.GetString("address")
	captcha_id:=this.GetString("captcha_id")
	captcha:=this.GetString("captcha")
	b:=cpt.Verify(captcha_id,captcha)
	if b != true {
		this.Data["json"]=responseResult("failure","Captcha Verify not match",20010)
	}else{
		result,msg,err:=models.CheckOrAddUser(username,password,email,tel,address)
		if result == "true"{
			this.Data["json"]=responseResult("success",msg,10000)
		}else if result == "false"{
			if msg=="existed" {
				this.Data["json"]=responseResult("failure",msg,20001)
			}else{
				this.Data["json"]=responseResult("failure",msg,20000)
			}
		}else if result == "error"{
			this.Data["json"]=responseResult("error",err,30000)
		}else{
			this.Data["json"]=responseResult("error","dont clear error",30010)
		}

	}
	this.ServeJSON()
	return

}
func (this *RegisterController) Put(){

}
func (this *RegisterController) Delete(){

}