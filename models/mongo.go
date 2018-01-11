package models

import (
	beego "github.com/beego/beego/v2/server/web"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Mongobase struct {
	query      *mgo.Query
	session    *mgo.Session
	collection *mgo.Collection
}

/**
 * 查询单条;
测试1:
    MongoId:=c.MongoId("54a6b7662b759fe4c0ff9b44")
    where := bson.M{"_id":MongoId}
    c.Dial("127.0.0.1","test","people")
    result:=c.One(where)
    c.dump(result)
测试2:
    result:=c.One(nil);
    c.dump(result)
*/
func (c *Mongobase) One(where interface{}) interface{} {
	var result interface{}
	c.collection.Find(where).One(&result)
	defer c.session.Close()
	return result
}

/**
 * 查询全部;
	where := bson.M{"name":"fengfeng"}
	c.Dial("127.0.0.1","test","people")
	result:=c.All(where)
	c.dump(result)
*/
func (c *Mongobase) All(where interface{}) interface{} {
	var result []interface{}
	c.collection.Find(where).All(&result)
	defer c.session.Close()
	return result
}

/**
 * 查询条数
测试1:
	result:=c.Count(nil)
测试2:
	where := bson.M{"name":"fengfeng"}
	result:=c.Count(where)
*/
func (c *Mongobase) Count(where interface{}) (int, error) {
	query := c.collection.Find(where)
	count, err := query.Count()
	defer c.session.Close()
	return int(count), err
}

/**
 * 执行自定义Mongo查询;
	测试:
	result:=c.Query(nil,1,10);
	result:=c.Query(nil,1,1);
*/
func (c *Mongobase) Query(where bson.M, skip int, limit int) interface{} {
	var result []interface{}
	c.collection.Find(where).Skip(skip).Limit(limit).Iter().All(&result)
	defer c.session.Close()
	return result
}

/**
 * 产生Mongo._id
	测试1:
	MongoId:=c.MongoId("54a6b7662b759fe4c0ff9b44")
	c.dump(MongoId)
*/
func (c *Mongobase) MongoId(id string) bson.ObjectId {
	return bson.ObjectIdHex(id)
}

/**
 * 插入数据
测试1:
person := models.Person{}
person.Use("person")
p:= models.Person{Name:"x"}
person.Insert(p);
*/
func (c *Mongobase) Insert(data interface{}) error {
	err := c.collection.Insert(data)
	defer c.session.Close()
	return err
}

/**
* 更新单条
   person := models.Person{}
   person.Use("person")
   p:= models.Person{Name:"xxx_xx"}
   MongoId:=person.MongoId("54a77af497479af50028a59f");
   where:=bson.M{"_id":MongoId}
   err:=person.Update(where,p);
   c.dump(err);
*/
func (c *Mongobase) Update(selector interface{}, update interface{}) error {
	err := c.collection.Update(selector, update)
	defer c.session.Close()
	return err
}

/**
* 更新单条,条件是_id的纯字符串;
   person := models.Person{}
   person.Use("person")
   p:= models.Person{Name:"xxxxxx"}
   err:=person.UpdateId("54a77af497479af50028a59f",p)
*/
func (c *Mongobase) UpdateId(id string, update interface{}) error {
	where := bson.M{"_id": c.MongoId(id)}
	err := c.collection.Update(where, update)
	defer c.session.Close()
	return err
}

/**
* 更新多条;
  测试:
   person := models.Person{}
   person.Use("person")
   p:= models.Person{Name:"xxx_xx"}
   MongoId:=person.MongoId("54a77af497479af50028a59f");
   where:=bson.M{"_id":MongoId}
   err:=person.UpdateAll(where,p);
   c.dump(err);
*/
func (c *Mongobase) UpdateAll(selector interface{}, update interface{}) error {
	_, err := c.collection.UpdateAll(selector, update)
	defer c.session.Close()
	return err
}

/**
* 更新,若没有,则插入
   测试:
   person := models.Person{}
   person.Use("person")
   p:= models.Person{Name:"xxx_xx"}
   MongoId:=person.MongoId("54a77af497479af50028a59f");
   where:=bson.M{"_id":MongoId}
   err:=person.Upsert(where,p);
   c.dump(err);
*/
func (c *Mongobase) Upsert(selector interface{}, update interface{}) error {
	_, err := c.collection.Upsert(selector, update)
	defer c.session.Close()
	return err
}

/**
 * 以conf为基础,连接mongo数据库
 */
func (c *Mongobase) Use(collect string) *Mongobase {
	//取配置;
	dsn, _ := beego.AppConfig.String("mongodsn")
	db, _ := beego.AppConfig.String("mongodb")

	//初始化;
	c.dial(dsn, db, collect)
	c.collection = c.session.DB(db).C(collect)
	return c
	// 这里不能提前关闭
	// defer c.session.Close()
}

/**
* 连接mongo数据库
  测试1:
  c.Dial("127.0.0.1","test","people")
  c.dump(c.collection)
*/
func (c *Mongobase) dial(ip string, db string, collect string) {
	c.session, _ = mgo.Dial(ip)
	c.session.SetMode(mgo.Monotonic, true)
	c.collection = c.session.DB(db).C(collect)
	// 这里不能提前关闭
	// defer c.session.Close()
}

/**
 * 删除单条;
 */
func (c *Mongobase) Remove(where interface{}) error {
	err := c.collection.Remove(where)
	defer c.session.Close()
	return err
}
