package main

import (
	"crmtest/controllers"
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	// index
	beego.Router("/", &controllers.IndexController{}, "get:Index")
	beego.Router("/file/Upload", &controllers.FileUploadController{}, "post:Upload")
	beego.Router("/file/download", &controllers.FileUploadController{}, "get:Download")
	beego.AutoRouter(&controllers.IndexController{})
	//beego.AutoRouter(&controllers.UserController{})
	beego.AutoRouter(&controllers.FileUploadController{})
	beego.SetStaticPath("/down", "upload/demo.txt")
	beego.AddFuncMap("hex", hex)
	beego.Run()
}

//mongoid
func hex(mid bson.ObjectId) string {
	return mid.Hex()
}
