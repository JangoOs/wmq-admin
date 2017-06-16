package main

import (
	_ "wmq-admin/app/routers"
	"github.com/astaxie/beego"
	"os"
	"wmq-admin/app/models"
)

func main() {
	//根据环境变量动态加载
	env := os.Getenv("GOENV");
	if(env == "") {
		env = "development";
	}
	if(env == "development") {
		beego.LoadAppConfig("ini", "conf/development.conf");
	}
	if(env == "testing") {
		beego.LoadAppConfig("ini", "conf/testing.conf");
	}
	if(env == "production") {
		beego.LoadAppConfig("ini", "conf/production.conf")
	}

	if beego.AppConfig.String("runmode") == "development" {
		beego.SetLevel(beego.LevelDebug)
	} else {
		beego.SetLevel(beego.LevelInformational)
		beego.SetLogger("file", `{"filename":"`+beego.AppConfig.String("log.log_file")+`"}`)
		beego.BeeLogger.DelLogger("console")
	}

	models.Init();
	beego.Run();
}