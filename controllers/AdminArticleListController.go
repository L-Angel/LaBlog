package controllers

import (
	"github.com/L-Angel/LaBlog/models"
	"github.com/astaxie/beego"
	"time"
	"strconv"
)

type  AdminArticleListController  struct{
	baseController
}

func (this *AdminArticleListController) Get(){
	this.Data["xsrf_token"]=this.XSRFToken()
	query:= map[string]string{}
	fields := make([]string,0)
	sortby:=[]string{"BlogId"}
	order := []string{"asc"}
	Blogs,err := models.GetAllBlogs(query,fields , sortby,order,0,10000)
	//beego.Info(Blogs)
	if err !=nil{
		beego.Error(err)
	}
	this.Data["Blogs"]=Blogs
	this.TplName="admin/AdminArticleList.html"
}
func (this *AdminArticleListController) Put(){
	release,boolerr:=this.GetBool("release")
	aid,err:=this.GetInt("aid")
	if boolerr !=nil{
		beego.Error("bool change error. "+boolerr.Error())
		this.Data["json"]=responseResult("error","",30012)
	}else if err!=nil {
		beego.Error("int change error. "+err.Error())
		this.Data["json"]=responseResult("error","",30013)
	}else if release == true{

		curTime:=time.Now().Local()
		article,err:=models.GetBlogsById(aid)
		if err !=nil{
			beego.Error(err)
			this.Data["json"]=responseResult("error","",30000)
		}else{
			article.IfRelease=1
			article.ReleaseTime=curTime
			models.UpdateBlogsById(article)
			this.Data["json"]=responseResult("true","",10000)
		}
	}else if release == false{
		curTime:=time.Now().Local()
		article,err:=models.GetBlogsById(aid)
		if err !=nil{
			beego.Error(err)
			this.Data["json"]=responseResult("error","",30000)
		}else{
			article.IfRelease=0
			article.ReleaseTime=curTime
			models.UpdateBlogsById(article)
			this.Data["json"]=responseResult("true","",10000)
		}
	}
	this.ServeJSON()

}

func (this *AdminArticleListController) Delete() {
	said:=this.Ctx.Input.Param(":aid")
	aid,err:=strconv.Atoi(said)
	if err!=nil{
		beego.Error(err)
		this.Data["json"]=responseResult("false",err.Error(),20010)
	}else {
		models.DeleteBlogs(aid);
		this.Data["json"]=responseResult("true","",10000)
	}

	/*
	aid :=this.GetString("aid")
	beego.Warn(aid)
	this.Data["json"]=responseResult("true","",1000)
	*/
	this.ServeJSON()
}