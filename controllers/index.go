/*
* @Author: Marte
* @Date:   2018-01-04 10:13:56
* @Last Modified by:   Marte
* @Last Modified time: 2018-01-05 14:48:26
 */

package controllers

import (
	. "crmtest/models"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"time"
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
func (this *IndexController) Index() {
	//this.display("index_index")
	this.display("file_upload")

}

//登陆执行
func (this *IndexController) Login() {
	fmt.Println("hello word")
	username := this.GetString("username")
	password := this.GetString("password")
	var m Mongobase
	where := bson.M{"username": username, "password": password}
	result := m.Use("user").One(where)
	if result != nil {
		user := result.(bson.M)
		uid := user["_id"].(bson.ObjectId)
		this.SetSession("username", username)
		this.SetSession("userid", uid.Hex())
		this.Redirect("/user/index", 302)
	} else {
		this.dump("用户名或密码错误！")
	}

}

//跳转注册页面
func (this *IndexController) Reg() {
	this.display("index_reg")
}

//注册接受
func (this *IndexController) Doreg() {
	username := this.GetString("username")
	password := this.GetString("password")
	var m Mongobase
	where := bson.M{"username": username}
	result := m.Use("user").One(where)
	if result != nil {
		this.dump("用户已存在")
	} else {
		u := &myuser{}
		u.Username = username
		u.Password = password
		m.Use("user").Insert(u)
		this.Redirect("/index/index", 302)

	}

}
