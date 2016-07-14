package controllers

import (
	"github.com/L-Angel/LaBlog/models"
	"github.com/L-Angel/LaBlog/tools"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/captcha"
	"github.com/astaxie/beego/cache"
)

var cpt *captcha.Captcha

func init(){
	store := cache.NewMemoryCache()
	cpt=captcha.NewWithFilter("/captcha/",store)
}

type LoginController struct  {
	baseController
}

func (this *LoginController) Get(){
	this.Data["IsLogin"] = true
	this.TplName="admin/Login.html"
	return
}

func (this *LoginController) Post(){
	username:=this.GetString("username")
	password:=tools.AES(this.GetString("password"))
	//password:=this.GetString("password")
	captcha_id:=this.GetString("captcha_id")
	captcha:=this.GetString("captcha")
	b:=cpt.Verify(captcha_id,captcha)
	if b != true {
		this.Data["json"]=responseResult("failure","Captcha Verify not match",20010)
	}else {
		user, err := models.GetUserByUsername(username)
		if err != nil {
			beego.Error(err)
			this.Data["json"] = responseResult("false", false, 20001)//无此用户
		} else if user.Password != password {
			beego.Warn("password error" + password)
			this.Data["json"] = responseResult("false", nil, 20010)//密码错误
		} else if user.Password == password {
			beego.Info("User athu successed!")
			this.SetSession("username",user.UserName)
			this.SetSession("auth",user.Authority)
			this.Data["json"] = responseResult("true", user, 10000)// 认证成功
		} else {
			beego.Warn("other error")
			this.Data["json"] = responseResult("error", nil, 30000)
		}
	}
	this.ServeJSON()


}

func (this *LoginController) Put(){

}
func (this *LoginController) Delete(){

}

