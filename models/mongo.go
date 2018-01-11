package models

import(
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    "github.com/astaxie/beego"
)

type Mongobase struct {
    query         *mgo.Query
    session       *mgo.Session
    collection    *mgo.Collection
}

/**
 * 查询单条;
测试1:
    MongoId:=this.MongoId("54a6b7662b759fe4c0ff9b44")
    where := bson.M{"_id":MongoId}
    this.Dial("127.0.0.1","test","people")
    result:=this.One(where)
    this.dump(result)
测试2:
    result:=this.One(nil);
    this.dump(result)
 */
func (this *Mongobase) One(where interface{}) interface{} {
    var result interface{}
    this.collection.Find(where).One(&result)
    defer this.session.Close()
    return result
}

/**
 * 查询全部;
	where := bson.M{"name":"fengfeng"}
	this.Dial("127.0.0.1","test","people")
	result:=this.All(where)
	this.dump(result)
 */
func (this *Mongobase) All(where interface{}) interface{} {
    var result [] interface{}
    this.collection.Find(where).All(&result)
    defer this.session.Close()
    return result
}

/**
 * 查询条数
测试1:
	result:=this.Count(nil)
测试2:
	where := bson.M{"name":"fengfeng"}
	result:=this.Count(where)
 */
func (this *Mongobase) Count(where interface{}) (int, error){
    query:=this.collection.Find(where)
    count,err:=query.Count()
    defer this.session.Close()
    return int(count),err
}

/**
 * 执行自定义Mongo查询;
	测试:
	result:=this.Query(nil,1,10);
	result:=this.Query(nil,1,1);
 */
func (this *Mongobase) Query(where bson.M,skip int,limit int) interface{} {
    var result [] interface{}
    this.collection.Find(where).Skip(skip).Limit(limit).Iter().All(&result)
    defer this.session.Close()
    return result
}

/**
 * 产生Mongo._id
	测试1:
	MongoId:=this.MongoId("54a6b7662b759fe4c0ff9b44")
	this.dump(MongoId)
 */
func (this *Mongobase) MongoId(id string) bson.ObjectId {
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
func (this *Mongobase) Insert(data interface{}) error{
    err := this.collection.Insert(data)
    defer this.session.Close();
    return err;
}

/**
 * 更新单条
    person := models.Person{}
    person.Use("person")
    p:= models.Person{Name:"xxx_xx"}
    MongoId:=person.MongoId("54a77af497479af50028a59f");
    where:=bson.M{"_id":MongoId}
    err:=person.Update(where,p);
    this.dump(err);
 */
func (this *Mongobase) Update(selector interface{}, update interface{}) error{
    err := this.collection.Update(selector,update)
    defer this.session.Close();
    return err
}

/**
 * 更新单条,条件是_id的纯字符串;
    person := models.Person{}
    person.Use("person")
    p:= models.Person{Name:"xxxxxx"}
    err:=person.UpdateId("54a77af497479af50028a59f",p) 
 */
func (this *Mongobase) UpdateId(id string, update interface{}) error{
    where := bson.M{"_id": this.MongoId(id)}
    err := this.collection.Update(where,update)
    defer this.session.Close();
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
    this.dump(err);
 */
func (this *Mongobase) UpdateAll(selector interface{}, update interface{}) error{
    _,err := this.collection.UpdateAll(selector,update)
    defer this.session.Close();
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
    this.dump(err);
 */
func (this *Mongobase) Upsert(selector interface{}, update interface{}) error{
    _,err := this.collection.Upsert(selector,update)
    defer this.session.Close();
    return err
}


/**
 * 以conf为基础,连接mongo数据库
 */
func (this *Mongobase) Use(collect string) *Mongobase {
    //取配置;
    dsn:=beego.AppConfig.String("mongodsn")
    db:=beego.AppConfig.String("mongodb")

    //初始化;
    this.dial(dsn,db,collect)
    this.collection =this.session.DB(db).C(collect)
    return this
    // 这里不能提前关闭
    // defer this.session.Close()
}

/**
 * 连接mongo数据库
   测试1:
   this.Dial("127.0.0.1","test","people")
   this.dump(this.collection)
 */
func (this *Mongobase) dial(ip string,db string,collect string){
    this.session,_ = mgo.Dial(ip)
    this.session.SetMode(mgo.Monotonic, true)
    this.collection =this.session.DB(db).C(collect)
    // 这里不能提前关闭
    // defer this.session.Close()
}

/**
 * 删除单条;
 */
func (this *Mongobase) Remove(where interface{}) error {
    err:=this.collection.Remove(where)
    defer this.session.Close()
    return err
}
