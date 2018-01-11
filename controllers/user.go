package controllers

import (
	. "crmtest/models"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type UserController struct {
	AuthController
}

type user struct {
	Id       bson.ObjectId `json:"id"        bson:"_id,omitempty"`
	Username string
	Password string
	Group    string
	Ctime    time.Time
}

//分类列表
func (this *UserController) Index() {
	var m Mongobase
	total, _ := m.Use("user").Count(nil)
	p, err := this.GetInt("page")
	if err != nil || p < 1 {
		p = 1
	}
	limit := 5
	skip := (p - 1) * limit
	result := m.Use("user").Query(nil, skip, limit)
	pager := NewPager(p, total, limit, "/user/index", true).ToString()
	this.Data["Result"] = result
	this.Data["Pager"] = pager
	this.display("user_index")
}

//添加分类
func (this *UserController) Add() {
	this.display("user_add")
}

//添加入库
func (this *UserController) Insert() {
	a := &user{}
	a.Username = this.GetString("username")
	a.Password = this.GetString("password")
	a.Group = this.GetString("group")
	a.Ctime = time.Now()
	var m Mongobase
	m.Use("user").Insert(a)
	this.Redirect("/user/index", 302)
}

//编辑入库
func (this *UserController) Edit() {
	var m Mongobase
	//http://127.0.0.1:8080/user/edit/?id=54c3be428c40e46373e69efe
	mid := m.MongoId(this.GetString("id"))
	where := bson.M{"_id": mid}
	result := m.Use("user").One(where)
	// this.dump(where)
	this.Data["result"] = result
	this.display("user_edit")
}

//添加入库
func (this *UserController) Update() {
	var m Mongobase
	//http://127.0.0.1:8080/user/edit/?id=54c3be428c40e46373e69efe
	mid := m.MongoId(this.GetString("id"))
	where := bson.M{"_id": mid}

	a := bson.M{"username": this.GetString("username"), "content": this.GetString("content"), "password": this.GetString("password"), "group": this.GetString("group")}
	update := bson.M{"$set": a}
	m.Use("user").Update(where, update)
	this.Redirect("/user/index", 302)
}

//删除
func (this *UserController) Delete() {
	var m Mongobase
	//http://127.0.0.1:8080/user/edit/?id=54c3be428c40e46373e69efe
	mid := m.MongoId(this.GetString("id"))
	where := bson.M{"_id": mid}
	m.Use("user").Remove(where)
	this.Redirect("/user/index", 302)
}
