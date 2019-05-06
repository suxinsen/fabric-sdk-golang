package main

import (
	"device-data-server/device"
	"device-data-server/fabric"
	_ "device-data-server/routers"
	"github.com/astaxie/beego"
	"log"
)

func init () {
	log.Println("********   init fabric sdk begin   ********")
	org         := beego.AppConfig.String("fabric.org")
	user        := beego.AppConfig.String("fabric.user")
	channel     := beego.AppConfig.String("fabric.channel")
	configFile  := beego.AppConfig.String("fabric.file")
	chainCodeId := beego.AppConfig.String("fabric.chaincodeid")

	sdkConfig := fabric.NewSdkConfig(channel, user, org, configFile, chainCodeId)
	sdkUtil := fabric.NewSdkUtil(sdkConfig)
	sdkUtil.Start()
	log.Println("********   init fabric sdk end     ********")
}

func main() {

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	go device.Start()
	beego.Run()
	defer fabric.ObtainSdkUtil().Sdk.Close()
}

