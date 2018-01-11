/*
* @Author: Marte
* @Date:   2018-01-04 10:13:56
* @Last Modified by:   Marte
* @Last Modified time: 2018-01-05 14:48:26
 */

package controllers

import (
	"fmt"
	"time"
	. "uploadfile/models"

	"gopkg.in/mgo.v2/bson"
)

type IndexController struct {
	BaseController
}
type myuser struct {
	Id       bson.ObjectId `json:"id"        bson:"_id,omitempty"`
	Username string
	Password string
	Group    string
	Ctime    time.Time
}

//登陆入口
func (c *IndexController) Index() {
	if c.GetSession("userid") != nil {
		c.Redirect("/file/index", 302)
	} else {
		c.display("index_index")
	}
}

//登陆执行
func (c *IndexController) Login() {
	fmt.Println("hello word")
	username := c.GetString("username")
	password := c.GetString("password")
	var m Mongobase
	where := bson.M{"username": username, "password": password}
	result := m.Use("user").One(where)
	if result != nil {
		user := result.(bson.M)
		uid := user["_id"].(bson.ObjectId)
		c.SetSession("username", username)
		c.SetSession("userid", uid.Hex())
		c.Redirect("/user/index", 302)
	} else {
		c.dump("用户名或密码错误！")
	}
}

//跳转注册页面
func (c *IndexController) Reg() {
	c.display("index_reg")
}

//注册接受
func (c *IndexController) Doreg() {
	username := c.GetString("username")
	password := c.GetString("password")
	var m Mongobase
	where := bson.M{"username": username}
	result := m.Use("user").One(where)
	if result != nil {
		c.dump("用户已存在")
	} else {
		u := &myuser{}
		u.Username = username
		u.Password = password
		m.Use("user").Insert(u)
		c.Redirect("/index/index", 302)

	}

}
