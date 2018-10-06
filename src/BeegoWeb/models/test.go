package models

import "github.com/astaxie/beego/orm"

type UserInfo struct {
	Id int64
	Username string
	Password string
}

func TestModel() int {
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/golang?charset=utf8", 30)
	orm.RegisterModel(new(UserInfo))  //自动映射为user_info表
	o := orm.NewOrm()
	user := UserInfo{Username:"aaa", Password:"6666"}
	id, _ := o.Insert(&user)
	return int(id)
}