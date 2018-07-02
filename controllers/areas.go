package controllers

import (
	"github.com/astaxie/beego"
	"Beego/Ihome/models"
	"github.com/astaxie/beego/orm"
)

type AreaController struct {
	beego.Controller
}
func (c *AreaController)GetArea(){
 beego.Info("GetArea 调用成功，调用API路由表")
 resp:=make(map[string]interface{})

 //
area:= []models.Area{}
o:=orm.NewOrm()
//err:=o.Read(&area)
num,err:=o.QueryTable("area").All(&area)
beego.Info("area data =",area)
//
if err!=nil{
	beego.Info("数据错误")
	resp["errno"]=4001
	resp["errmsg"]="查询失败"
	c.RetData(resp)
	return
}
if num==0{
	resp["errno"]=models.RECODE_NODATA
	resp["errmsg"]=models.RecodeText(models.RECODE_NODATA)
	c.RetData(resp)
	return
}
//
resp["errno"]=models.RECODE_OK
resp["errmsg"]=models.RecodeText(models.RECODE_OK)
resp["data"]=area
beego.Info(resp)
c.RetData(resp)
}
func (this *AreaController) RetData(resp map[string]interface{}){
	this.Data["json"]=resp
	this.ServeJSON()
}