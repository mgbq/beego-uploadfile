package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"os"
)

type BaseController struct {
	beego.Controller
	moduleName     string
	controllerName string
	actionName     string
	options        map[string]string
}

func (this *BaseController) dump(params interface{}) {
	fmt.Fprint(this.Ctx.ResponseWriter,params);
}

//取cookie
func (this *BaseController) GetCookie(key string)interface{}{
	return this.Ctx.GetCookie(key)
}

//存cookie
func (this *BaseController) SetCookie(key string,value string,time int32){
	this.Ctx.SetCookie(key,value,time)
	return
}

//框架缓存;
func (this *BaseController) NewCache() cache.Cache {
	// c,_:=cache.NewCache("file", `{"CachePath":"./cache"}`)
	c,_:=cache.NewCache("memory", `{"interval":60}`)
	return c
}

//显示页面;
func (this *BaseController) display(tpl string) {
	var theme string
	if v, ok := this.options["theme"]; ok && v != "" {
		theme = v
	} else {
		theme = "default"
	}
	if _, err := os.Stat(beego.BConfig.WebConfig.ViewsPath + "/" + theme + "/layout.html"); err == nil {
		this.Layout = theme + "/layout.html"
	}
	this.TplName = theme + "/" + tpl + ".html"
}
