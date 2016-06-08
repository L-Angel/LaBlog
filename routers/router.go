package routers

import (
	"github.com/L-Angel/lablog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.ErrorController(&controllers.ErrorController{})
        beego.Router("*.md",   &controllers.MarkdownController{})
	beego.Router("/", &controllers.MainController{})
	beego.Router("/links",&controllers.LinksController{})
	beego.Router("/share",&controllers.ShareController{})
	beego.Router("/category",&controllers.CategoryController{})
	beego.Router("/about",&controllers.AboutController{})

	beego.Router("/admin/login",&controllers.LoginController{})
	beego.Router("/admin/register",&controllers.RegisterController{})
}
