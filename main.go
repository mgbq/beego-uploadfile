package main

import (
	_ "uploadfile/routers"

	beego "github.com/beego/beego/v2/server/web"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	beego.AddFuncMap("hex", hex)
	beego.Run()
}

//mongoid
func hex(mid bson.ObjectId) string {
	return mid.Hex()
}
