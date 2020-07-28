package dao

import (
	"fmt"
	"log"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/vhaoran/vchat/common/ypage"
	"github.com/vhaoran/vchat/common/yverify"
	"github.com/vhaoran/vchat/lib/ylog"
	"github.com/vhaoran/vchat/lib/ypg"
)

type (
	Abc struct {
		Id   int64
		Name string
		Age  string
	}

	AbcDao struct {
	}
)

func (d *AbcDao) Get(id string) (Abc, error) {
	//	通过id获取记录
	var a1 Abc
	ypg.X.First(&a1, "id=?", id)
	ylog.Info("查询到的数据是", a1)
	return a1, nil
}
func (d *AbcDao) Exist(id string) bool {
	//	通过id获取记录
	var a1 Abc
	ret := ypg.X.First(&a1, "id=?", id)
	if ret.RowsAffected > 0 {
		return true
	}
	return false

}

var db = newDB()

func newDB() *gorm.DB {
	connStr := fmt.Sprintf("host=%s port=5432 user=%s dbname=%s password=%s sslmode=disable", "localhost", "postgres", "test", "123456")
	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		log.Println(err)
	}
	db.DB().SetMaxOpenConns(500)
	db.LogMode(true)
	return db
}
func (d *AbcDao) ListAll() ([]*Abc, error) {
	//查询所有记录
	abs := make([]*Abc, 0)
	//Find是查找所有的记录
	db.Find(&abs)
	return abs, nil
}
func (d *AbcDao) GetAll() ([]*Abc, error) {
	l := make([]*Abc, 0)
	ret := ypg.X.Find(&l)
	ylog.Debug("操作影响的行数", ret.RowsAffected)
	return l, nil

}
func (d *AbcDao) GetByName(name string) ([]*Abc, error) {
	l := make([]*Abc, 0)
	ypg.X.Find(&l, "name=?", name)
	return l, nil
}
func (d *AbcDao) Insert(bean *Abc) (int64, error) {
	ret := ypg.X.Save(bean)
	ylog.Info("执行插入操作后影响的行数是", ret.RowsAffected)
	return 0, nil
}
func (d *AbcDao) Rm(id string) error {
	id1, _ := strconv.ParseInt(id, 10, 64)
	bean := &Abc{Id: id1}
	ret := ypg.X.Delete(bean)
	ylog.Info("执行删除操作后影响的行数是", ret.RowsAffected)
	return nil
}

func (d *AbcDao) Ch(bean *Abc) error {
	ret := ypg.X.Save(bean)
	ylog.Info("影响的行数", ret.RowsAffected)
	return nil
	//还没想好该返回什么

}
func (d *AbcDao) Page(pb *ypage.PageBeanMap) (*ypage.PageBeanMap, error) {
	if err := yverify.NewObj().
		NotNilPtr("传入参数不能为空", pb).
		Err(); err != nil {
		return nil, nil
	}
	l := make([]*Abc, 0)
	ret, err := ypg.PageMap(ypg.X, &l, pb)
	if err != nil {
		ylog.Error(err)
	}
	ylog.Debug(ret)
	return ret, nil
}
