package main

import (
	_ "Beego/Ihome/routers"
	_ "Beego/Ihome/models"
	_ "github.com/astaxie/beego/session/mysql"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"strings"
	"net/http"
)

func main() {
	ignoreStaticPath()
	beego.Run()
}

func ignoreStaticPath(){
	beego.InsertFilter("/",beego.BeforeRouter,TransparentStatic)
	beego.InsertFilter("/*",beego.BeforeRouter,TransparentStatic)

}

func TransparentStatic(ctx *context.Context){
	orpath:=ctx.Request.URL.Path
	beego.Debug("request url:",orpath)
	//检测API 字段
	if strings.Index(orpath,"api")>=0{
		return
	}
	http.ServeFile(ctx.ResponseWriter,ctx.Request,"static/html/"+ctx.Request.URL.Path)
}