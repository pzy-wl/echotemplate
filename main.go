package main

import (
	"html/template"
	"io"

	echo "github.com/labstack/echo/v4"

	"template/intf"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
func main() {
	e := echo.New()
	tpl := &Template{
		//这是扫描有关的html模板
		templates: template.Must(template.ParseGlob("template/*.html")),
	}
	e.Renderer = tpl
	e.Any("/", intf.Index)
	e.GET("/get", intf.Get)
	e.GET("/show", intf.Show)
	e.GET("/query", intf.Query)
	e.GET("/rm", intf.Rm)
	e.GET("/add", intf.Add)
	e.GET("/ch", intf.Ch)
	e.POST("/getById", intf.GetById)
	e.POST("/deleteById", intf.DeleteById)
	e.POST("/insert", intf.Insert)
	e.POST("/update", intf.Update)
	e.Logger.Fatal(e.Start(":2020"))
}
