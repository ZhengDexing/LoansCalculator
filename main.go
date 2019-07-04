package main

import (
	_ "LoansCalculator/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.SetStaticPath("/", "views")
	beego.Run()
}
