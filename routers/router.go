package routers

import (
	"device-data-server/controllers"
	"github.com/astaxie/beego"
	"log"
)

func init() {
	log.Print("******   init router   ******")
	beego.Router("/device/obtain", &controllers.DeviceController{}, "get:Obtain")
	beego.Router("/device/upload", &controllers.DeviceController{}, "post:Upload")
}
