package controllers

import (
	"fmt"
	"os"

	"github.com/beego/beego/v2/client/cache"
	beego "github.com/beego/beego/v2/server/web"
)

type BaseController struct {
	beego.Controller
	moduleName     string
	controllerName string
	actionName     string
	options        map[string]string
}

func (c *BaseController) dump(params interface{}) {
	fmt.Fprint(c.Ctx.ResponseWriter, params)
}

//取cookie
func (c *BaseController) GetCookie(key string) interface{} {
	return c.Ctx.GetCookie(key)
}

//存cookie
func (c *BaseController) SetCookie(key string, value string, time int32) {
	c.Ctx.SetCookie(key, value, time)
	return
}

//框架缓存;
func (c *BaseController) NewCache() cache.Cache {
	// c,_:=cache.NewCache("file", `{"CachePath":"./cache"}`)
	ca, _ := cache.NewCache("memory", `{"interval":60}`)
	return ca
}

//显示页面;
func (c *BaseController) display(tpl string) {
	var theme string
	if v, ok := c.options["theme"]; ok && v != "" {
		theme = v
	} else {
		theme = "default"
	}
	if _, err := os.Stat(beego.BConfig.WebConfig.ViewsPath + "/" + theme + "/layout.html"); err == nil {
		c.Layout = theme + "/layout.html"
	}
	c.TplName = theme + "/" + tpl + ".html"
}
