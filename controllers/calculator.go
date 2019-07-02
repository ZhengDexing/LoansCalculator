package controllers

import (
	"LoansCalculator/entity"
	"LoansCalculator/service"
	"LoansCalculator/util"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

type Calculator struct {
	beego.Controller
}

func (c *Calculator) Post() {
	var calculatorInput entity.CalculatorInput
	// 获取requestBody
	data := c.Ctx.Input.RequestBody
	// json解析成对象
	err := json.Unmarshal(data, &calculatorInput)
	if err != nil {
		result := util.Result(util.ERROR, nil, "请求解析失败！")
		c.Data["json"] = result
		c.ServeJSON()
		return
	}

	fmt.Println("calculatorInput:", calculatorInput)

	var outPut entity.CalculatorOutput

	switch calculatorInput.PaymentMethod {
	case 0:
		outPut = service.InterestAvg(calculatorInput)
	case 1:
		outPut = service.MoneyAvg(calculatorInput)
	default:
		c.Data["json"] = util.Result(util.ERROR, nil, "请求解析失败！")
		c.ServeJSON()
		return
	}

	result := util.Result(util.SUCCESS, outPut, "success")
	c.Data["json"] = result
	c.ServeJSON()
}
