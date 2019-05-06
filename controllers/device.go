package controllers

import (
	"device-data-server/fabric"
	"github.com/astaxie/beego"
	"log"
)

type DeviceController struct {
	beego.Controller
}

func (d *DeviceController) Upload() {
	code := d.GetString("code")
	data := d.GetString("data")
	if len(code) == 0 || len(data) == 0 {
		d.Data["json"] = "Invalid args, system error."
	} else {
		log.Printf("{code:%s,data:%s}", code, data)
		//upload data to fabric network
		var args [][]byte
		args = append(args, []byte(code))
		args = append(args, []byte(data))
		result := fabric.ObtainSdkUtil().Invoke("upload", args)
		d.Data["json"] = result
	}
	d.ServeJSON()
}

func (d *DeviceController) Obtain() {
	code := d.GetString("code")
	if len(code) == 0 {
		d.Data["json"] = "Invalid args, system error."
	} else {
		log.Printf("obtain code: %s", code)
		//obtain data to fabric network
		var args [][]byte
		args = append(args, []byte(code))
		result := fabric.ObtainSdkUtil().Obtain("obtain", args)
		d.Data["json"] = result
	}
	d.ServeJSON()
}

