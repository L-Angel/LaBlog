package controllers

import (
	"time"
	"io/ioutil"
	"github.com/astaxie/beego"
	"os"
)
var directory = beego.AppConfig.String("blogfiledirectory")

type  MarkdownEditorController struct  {
	baseController
}

func (this *MarkdownEditorController) Get(){
	this.Data["xsrf_token"]=this.XSRFToken()
	this.TplName="admin/MarkdownEditor.html"
}

func (this *MarkdownEditorController) Post(){
	//articleId:=this.GetInt("articleId")
	//articleUniqId:=this.GetSteing("articleUniqId")
	//articleTitle:=this.GetString("articleTitle")
	//categoryId:=this.GetInt("catetoryId")
	//author:=this.GetString("author")
	articleContent:=this.GetString("articleContent")
	beego.Warn(articleContent)
	curTime:=time.Now().Local()
	articleUniqId := curTime.Format("20060102150405")
        if ok := saveFile(directory,articleUniqId,articleContent); ok == true{
		beego.Info(articleUniqId+".md save successed!")
		this.Data["json"]=responseResult("true",nil,10000)
        }else{
		this.Data["json"]=responseResult("false",nil,20000)
	}
	this.ServeJSON()
}

func saveFile(directory,name,content string)(bool){
	path:=directory+"/"+name+".md"
	beego.Info(content)
	err := ioutil.WriteFile( path, []byte(content), os.ModePerm )
	if err != nil{
		beego.Error(err)
		return false
	}else{
		beego.Info(path)
		return true
	}
}

func getFile(directory,name string)(string,bool){
	path:=directory+"/"+name+".md"
	data,err:=ioutil.ReadFile(path)
	if err !=nil{
		beego.Error(err)
		return "",false
	}else{
		beego.Info(path)
		return string(data),true
	}

}

func responseResult(r string,m interface{},statusCode int)( map[string]interface{}){
	return map[string]interface{}{
		"result":r,
		"msg":m,
		"statuscode" : statusCode,
	}
}
