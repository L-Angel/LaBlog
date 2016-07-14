package main

import (
	_ "github.com/L-Angel/lablog/routers"
	"github.com/astaxie/beego"
	"github.com/L-Angel/LaBlog/models"
)

func main() {
	models.RegisterDataBase()
	beego.Run()
}

