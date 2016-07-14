package controllers

import (
	"github.com/L-Angel/LaBlog/models"
	"github.com/astaxie/beego"
	"time"
	"strconv"
)

type AdminModifyArticleController struct {
	baseController
}

func (this *AdminModifyArticleController) Get(){
	//XSRF
	this.Data["xsrf_token"]=this.XSRFToken()
	//获取文章内容
	articleid,err:=this.GetInt("aid")
	if err !=nil {
		beego.Error(err)
	}
	/*
	id,strconverr:=strconv.Atoi(articleid)
	if strconverr != nil{
		beego.Error(strconverr)
	}
	beego.Info(id)
	*/
	beego.Info(articleid)
	BlogContent,err:=models.GetBlogsById(articleid)
	if err!=nil{
		beego.Error(err)
	}else{
		this.Data["Blog"]=BlogContent
	}
	//查询文章分类列表
	query:=map[string]string{}
	fileds:=[]string{}
	sortby:=[]string{"CategoryId"}
	order:=[]string{"asc"}
	categorys,err := models.GetAllCategory(query,fileds,sortby,order,0,10000)
	if err !=nil{
		beego.Error(err)
	}else{
		this.Data["categorys"]=categorys
	}
        //获取文章
	 if articlemdcontent,ok:=getFile(directory,BlogContent.BlogUniqId); ok==true{
		 this.Data["content"]=articlemdcontent
	 }else {
		 beego.Error("get Blog file Error!")
	 }
	//选择静态引入文件
	this.Data["modify"]=true
	this.TplName="admin/MarkdownEditor.html"
}

func (this *AdminModifyArticleController) Post(){
	articleTitle:=this.GetString("articleTitle")
	categoryId,_:=this.GetInt("categoryId")
	author:=this.GetString("author")
	articleContent:=this.GetString("articleContent")
	summary:=this.GetString("summary")
	articleUniqId:=this.GetString("uniqid")
	articleUniqId=PureDigitString(articleUniqId)
	aid:=this.GetString("articleid")
	aid=PureDigitString(aid)
	articleid,err:=strconv.Atoi(aid)
	if err !=nil{
		beego.Error(err)
	}
	beego.Info(articleUniqId)
	if ok := saveFile(directory,articleUniqId,articleContent); ok == true{
		blog:= models.Blogs{Id:articleid}
		curTime:=time.Now().Local()
		blog.Author=author
		blog.BlogTitle=articleTitle
		category,e:=models.GetCategoryById(categoryId)
		blog.CategoryCategoryId=category
		blog.ReleaseTime=curTime
		blog.Summary=summary
		blog.BlogUniqId=articleUniqId
		blog.FilePath=directory+"/"+articleUniqId+".md"
		err:=models.UpdateBlogsById(&blog)
		if e !=nil {
			beego.Error(e)
			this.Data["json"]=responseResult("false","fk error",20010)
		}else if err!=nil{
			beego.Error(err)
			this.Data["json"]=responseResult("false","db error",20020)
		}else if e==nil && err ==nil{
			beego.Info(articleUniqId+".md save successed!")
			this.Data["json"]=responseResult("true","success",10000)
		}
	}else{
		beego.Error("Blog file save error!")
		this.Data["json"]=responseResult("false",nil,20000)
	}
	this.ServeJSON()
}