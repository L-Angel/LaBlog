package controllers

import (
	"io/ioutil"
	"github.com/astaxie/beego"
	"os"
	"regexp"
)
func responseResult(r string,m interface{},statusCode int)( map[string]interface{}){
	return map[string]interface{}{
		"result":r,
		"msg":m,
		"statuscode" : statusCode,
	}
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

func PureDigitString(s string)(newstr string){
	reg:=regexp.MustCompile("[0-9]+")
	newstr=reg.FindString(s)
	return
}
