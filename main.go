package main

import (
	"github.com/PPGo_ApiAdmin/models"
	_ "github.com/PPGo_ApiAdmin/routers"

	"github.com/astaxie/beego"
)

func main() {
	models.Init()
	beego.Run()
}
