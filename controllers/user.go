package controllers

import (
	"time"
	. "uploadfile/models"

	"gopkg.in/mgo.v2/bson"
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
func (c *UserController) Index() {
	var m Mongobase
	total, _ := m.Use("user").Count(nil)
	p, err := c.GetInt("page")
	if err != nil || p < 1 {
		p = 1
	}
	limit := 5
	skip := (p - 1) * limit
	result := m.Use("user").Query(nil, skip, limit)
	pager := NewPager(p, total, limit, "/user/index", true).ToString()
	c.Data["Result"] = result
	c.Data["Pager"] = pager
	c.display("user_index")
}

//添加分类
func (c *UserController) Add() {
	c.display("user_add")
}

//添加入库
func (c *UserController) Insert() {
	a := &user{}
	a.Username = c.GetString("username")
	a.Password = c.GetString("password")
	a.Group = c.GetString("group")
	a.Ctime = time.Now()
	var m Mongobase
	m.Use("user").Insert(a)
	c.Redirect("/user/index", 302)
}

//编辑入库
func (c *UserController) Edit() {
	var m Mongobase
	//http://127.0.0.1:8080/user/edit/?id=54c3be428c40e46373e69efe
	mid := m.MongoId(c.GetString("id"))
	where := bson.M{"_id": mid}
	result := m.Use("user").One(where)
	// c.dump(where)
	c.Data["result"] = result
	c.display("user_edit")
}

//添加入库
func (c *UserController) Update() {
	var m Mongobase
	//http://127.0.0.1:8080/user/edit/?id=54c3be428c40e46373e69efe
	mid := m.MongoId(c.GetString("id"))
	where := bson.M{"_id": mid}

	a := bson.M{"username": c.GetString("username"), "content": c.GetString("content"), "password": c.GetString("password"), "group": c.GetString("group")}
	update := bson.M{"$set": a}
	m.Use("user").Update(where, update)
	c.Redirect("/user/index", 302)
}

//删除
func (c *UserController) Delete() {
	var m Mongobase
	//http://127.0.0.1:8080/user/edit/?id=54c3be428c40e46373e69efe
	mid := m.MongoId(c.GetString("id"))
	where := bson.M{"_id": mid}
	m.Use("user").Remove(where)
	c.Redirect("/user/index", 302)
}
