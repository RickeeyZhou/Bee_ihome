package controllers

import (
	"github.com/astaxie/beego"
	"Beego/Ihome/models"
)

type SessionController struct {
	beego.Controller
}
func (this *SessionController) RetData(resp map[string]interface{})  {
	this.Data["json"] = resp
	this.ServeJSON()
}
func (this *SessionController)GetSessionData(){
	resp:=make(map[string]interface{})
	defer this.RetData(resp)
	user:=models.User{}
	//user.Name="ss"
	resp["errno"]=models.RECODE_DBERR
	resp["errmsg"]=models.RecodeText(models.RECODE_DBERR)
	//resp["data"]=user
	name:=this.GetSession("name")
	if name!=nil{
		user.Name= name.(string)
		resp["errno"]=models.RECODE_OK
		resp["errmsg"]=models.RecodeText(models.RECODE_OK)
		resp["data"]=user
	}
	//this.RetData(resp)
}
func (this *SessionController)DeleteSession(){
	resp:=make (map[string]interface{})
	defer this.RetData(resp)
	this.DelSession("name")
	resp["errno"]=models.RECODE_OK
	resp["errmsg"]=models.RecodeText(models.RECODE_OK)

}