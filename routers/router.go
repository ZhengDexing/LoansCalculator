package routers

import (
	"LoansCalculator/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/calculator", &controllers.Calculator{})
}
