package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"github.com/astaxie/beego/orm"
	"Beego/Ihome/models"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) RetData(resp map[string]interface{})  {
	this.Data["json"] = resp
	this.ServeJSON()
}

func (this *UserController) Reg() {
	resp:=make(map[string]interface{})
	defer this.RetData(resp)
	json.Unmarshal(this.Ctx.Input.RequestBody,&resp)

	beego.Info("111111111111111-----",resp)
	beego.Info(`resp["mobile"] = `,resp["mobile"])
	beego.Info(`resp["password"] = `,resp["password"])
	beego.Info(`resp["sms_code"] =`,resp["sms_code"])
	o:=orm.NewOrm()
	user:=models.User{}
	user.Password_hash=resp["password"].(string)
	user.Name=resp["mobile"].(string)
	user.Mobile=resp["mobile"].(string)
	//查u数据库
	id ,err:=o.Insert(&user)
	if err!=nil{
		resp["errno"]=models.RECODE_NODATA
		resp["errmsg"]=models.RecodeText(models.RECODE_NODATA)
		return
	}
	resp["errno"]=models.RECODE_OK
	resp["errmsg"]=models.RecodeText(models.RECODE_OK)
	beego.Info("reg succee ,id =",id)
	this.SetSession("name",user.Name)



}