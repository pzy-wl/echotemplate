package intf

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/vhaoran/vchat/lib"

	"template/dao"
)

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

var AbaDaox dao.AbcDao

func Index(e echo.Context) error {
	return e.Render(http.StatusOK, "index.html", "")
}
func Get(e echo.Context) error {
	id := e.FormValue("id")
	name := e.FormValue("name")
	return e.String(http.StatusOK, "参数是id:"+id+"name:"+name)
}
func Show(c echo.Context) error {
	//如果是定义的,则直接引用参数名,如果是文件,则需要加上后缀名".html"
	//l:=[]string{"hello", "world", "this","is", "echo"}
	id := c.FormValue("id")
	println("id是:", id)
	l, _ := AbaDaox.GetAll()
	return c.Render(http.StatusOK, "show.html", l)
}
func GetById(e echo.Context) error {
	println("-----------GetById---------")
	id := e.FormValue("id")
	bean, _ := AbaDaox.Get(id)
	return e.Render(http.StatusOK, "res.html", bean)
}
func Query(e echo.Context) error {
	return e.Render(http.StatusOK, "query.html", "请输入用户id")
}
func Rm(e echo.Context) error {
	return e.Render(http.StatusOK, "del.html", "请输入需要删除的用户id")
}
func Add(e echo.Context) error {
	return e.Render(http.StatusOK, "add.html", "")
}
func Ch(e echo.Context) error {
	return e.Render(http.StatusOK, "ch.html", "")
}
func Update(e echo.Context) error {
	bean := new(dao.Abc)
	bean.Id, _ = strconv.ParseInt(e.FormValue("id"), 10, 64)
	if AbaDaox.Exist(strconv.FormatInt(bean.Id, 10)) {
		bean.Name = e.FormValue("name")
		bean.Age = e.FormValue("age")
		AbaDaox.Ch(bean)
		return e.Render(http.StatusOK, "res.html", "用户信息修改成功")

	}
	return e.Render(http.StatusBadRequest, "res.html", "修改用户信息失败,请检查用户是否存在!")

}
func Insert(e echo.Context) error {
	bean := new(dao.Abc)
	bean.Name = e.FormValue("name")
	bean.Age = e.FormValue("age")
	ret, _ := AbaDaox.Insert(bean)
	if ret > 0 {
		return e.Render(http.StatusOK, "res.html", "插入数据成功!")
	}
	return e.Render(http.StatusOK, "res.html", "添加用户失败!")

}
func DeleteById(e echo.Context) error {
	id := e.FormValue("id")
	if AbaDaox.Exist(id) {
		AbaDaox.Rm(id)
		return e.Render(http.StatusOK, "res.html", "删除成功!")
	} else {
		return e.Render(http.StatusOK, "res.html", "不存在此id的用户")
	}

}
