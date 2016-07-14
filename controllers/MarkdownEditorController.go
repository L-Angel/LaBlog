package controllers

import (
	"time"
	"github.com/astaxie/beego"
	"github.com/L-Angel/LaBlog/models"
)
var directory = beego.AppConfig.String("blogfiledirectory")

type  MarkdownEditorController struct  {
	baseController
}

func (this *MarkdownEditorController) Get(){
	//xsrf
	this.Data["xsrf_token"]=this.XSRFToken()
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
	this.TplName="admin/MarkdownEditor.html"
}

func (this *MarkdownEditorController) Post(){
	//articleId:=this.GetInt("articleId")
	//articleUniqId:=this.GetSteing("articleUniqId")
	articleTitle:=this.GetString("articleTitle")
	categoryId,_:=this.GetInt("categoryId")
	author:=this.GetString("author")
	articleContent:=this.GetString("articleContent")
	summary:=this.GetString("summary")
	curTime:=time.Now().Local()
	articleUniqId := curTime.Format("20060102150405")
        if ok := saveFile(directory,articleUniqId,articleContent); ok == true{
		blog:=new(models.Blogs)
		blog.Author=author
		blog.BlogTitle=articleTitle
		category,e:=models.GetCategoryById(categoryId)
		blog.CategoryCategoryId=category
		blog.CreateTime=curTime
		blog.ReleaseTime=curTime
		blog.Summary=summary
		blog.BlogUniqId=articleUniqId
		blog.FilePath=directory+"/"+articleUniqId+".md"
		id,err:=models.AddBlogs(blog)
		if e !=nil {
			beego.Error(e)
			this.Data["json"]=responseResult("false","fk error",20010)
		}else if err!=nil{
			beego.Error(err)
			this.Data["json"]=responseResult("false","db error",20020)
		}else if e==nil && err ==nil{
			beego.Info(articleUniqId+".md save successed!")
			msg:=map[string]interface{}{"id":id}
			this.Data["json"]=responseResult("true",msg,10000)
		}

        }else{
		beego.Error("Blog file save error!")
		this.Data["json"]=responseResult("false",nil,20000)
	}
	this.ServeJSON()
}
func (this *MarkdownEditorController) Delete(){

}

func (this *MarkdownEditorController) Put(){

}
