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
	beego.Router("/resume",&controllers.ResumeController{})


	beego.Router("/admin/login",&controllers.LoginController{})
	beego.Router("/admin/register",&controllers.RegisterController{})
	beego.Router("/admin/index",&controllers.AdminIndexController{})
	beego.Router("/admin/editor",&controllers.MarkdownEditorController{})
	beego.Router("/admin/articlecategory",&controllers.AdminArticleCategoryController{})
	beego.Router("/admin/articlelist",&controllers.AdminArticleListController{})
	beego.Router("/admin/userlist",&controllers.AdminUserListController{})
	beego.Router("/admin/adduser",&controllers.AdminAddUserController{})
	beego.Router("/admin/pictures",&controllers.AdminPicturesController{})
	beego.Router("/admin/addpicture",&controllers.AdminAddPictureController{})

}
