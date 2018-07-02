package controllers

import (
	"github.com/astaxie/beego"
	"Beego/Ihome/models"
)

type HouseIndexController struct {
	beego.Controller
}
func (this *HouseIndexController)ReData(resp map[string]interface{}){
	this.Data["json"]=resp
	this.ServeJSON()
}
func (this *HouseIndexController)GetHouseIndex(){
	resp:=make(map[string]interface{})
	resp["errno"]=models.RECODE_DBERR
	resp["errmsg"]=models.RecodeText(models.RECODE_DBERR)
	this.ReData(resp)
}