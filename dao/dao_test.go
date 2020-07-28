package dao

import (
	"fmt"
	"testing"

	"github.com/vhaoran/vchat/lib"
	"github.com/vhaoran/vchat/lib/ylog"
)

var AbcDaox = new(AbcDao)

func init() {
	_, err := lib.InitModulesOfOptions(&lib.LoadOption{
		LoadMicroService: false,
		LoadEtcd:         false,
		LoadPg:           true,
		LoadRedis:        true,
		LoadMongo:        false,
		LoadMq:           false,
		LoadRabbitMq:     false,
		LoadJwt:          false,
	})
	if err != nil {
		panic(err.Error())
	}
}
func Test_get(t *testing.T) {
	a, err := AbcDaox.Get("13")
	if err != nil {
		ylog.Error("Testing get------------>", err)
	}
	ylog.Info(a)
}
func Test_getAll(t *testing.T) {

	l, _ := AbcDaox.GetAll()
	//l, _:=AbcDaox.ListAll()
	for _, i := range l {
		fmt.Println(i)
	}
}
func Test_insert(t *testing.T) {
	bean := &Abc{
		Name: "lmm",
		Age:  "1312",
	}
	AbcDaox.Insert(bean)
}
func Test_update(t *testing.T) {
	bean1 := &Abc{
		Id:   15,
		Name: "liu",
		Age:  "13",
	}
	AbcDaox.Ch(bean1)
}
