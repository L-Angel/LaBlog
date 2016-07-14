package controllers

import (
	"github.com/astaxie/beego"
	"github.com/L-Angel/LaBlog/models"
	"github.com/astaxie/beego/context"
)

type MainController struct {
	baseController
}

func (this *MainController) Get() {
	query:= map[string]string{}
	fields := make([]string,0)
	sortby:=[]string{"BlogId"}
	order := []string{"asc"}
	Blogs,err := models.GetAllBlogs(query,fields , sortby,order,0,10000)
	beego.Info(Blogs)
	beego.Info(err)
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	if err != nil{
		beego.Error(err)
		this.Abort("500")
		this.EnableRender=false
	}else{
		this.Data["Blogs"]=Blogs
		this.TplName = "index.html"
	}
	this.TplName = "index.html"
}


var FilterUser = func(ctx *context.Context){
	if ctx.Request.RequestURI != "/admin/login"{
		ctx.Redirect(302,"/admin/login")
	}

	/*
	val,ok:=ctx.Input.Session("login").(string)

	if !ok||val==""{
		if ctx.Request.Method == "GET"{
			ctx.Redirect(302,"/admin/login")
		}else if ctx.Request.Method == "POST"{
			ctx.Redirect(302,"/admin/login")
		}
	}
	*/
}