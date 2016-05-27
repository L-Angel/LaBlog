package routers

import (
	"github.com/L-Angel/lablog/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("*.md",   &controllers.MarkdownController{})
    beego.Router("/", &controllers.MainController{})
}
