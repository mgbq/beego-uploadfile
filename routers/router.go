package routers

import (
	"uploadfile/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	//beego.Router("/", &controllers.IndexController{}, "get:Index")
	//beego.Router("/file/Upload", &controllers.FileController{}, "POST:Upload")
	// index
	beego.Router("/", &controllers.IndexController{}, "get:Index")
	beego.Router("/file/Upload", &controllers.FileUploadController{}, "post:Upload")
	beego.Router("/file/download", &controllers.FileUploadController{}, "get:Download")
	beego.Router("/file/index", &controllers.FileUploadController{}, "get:Index")
	beego.AutoRouter(&controllers.IndexController{})
	beego.AutoRouter(&controllers.UserController{})
	beego.AutoRouter(&controllers.FileUploadController{})
	beego.SetStaticPath("/down", "upload/demo.txt")
}
