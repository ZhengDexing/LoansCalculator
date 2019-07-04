package controllers

import (
	"LoansCalculator/entity"
	"LoansCalculator/service"
	"LoansCalculator/util"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
)

type Calculator struct {
	beego.Controller
}

func (c *Calculator) Post() {
	var calculatorInput entity.CalculatorInputStr
	// 获取requestBody
	data := c.Ctx.Input.RequestBody
	// json解析成对象
	err := json.Unmarshal(data, &calculatorInput)
	if err != nil {
		fmt.Println(err)
		result := util.Result(util.ERROR, nil, "请求解析失败！")
		c.Data["json"] = result
		c.ServeJSON()
		return
	}

	fmt.Println("calculatorInput:", calculatorInput)

	var outPut entity.Result
	investment, err := strconv.ParseFloat(calculatorInput.Investment, 64)
	proportion, err := strconv.ParseFloat(calculatorInput.Proportion, 64)
	years, err := strconv.Atoi(calculatorInput.Years)
	interestRate, err := strconv.ParseFloat(calculatorInput.InterestRate, 64)
	upProportion, err := strconv.ParseFloat(calculatorInput.UpProportion, 64)
	paymentMethod, err := strconv.Atoi(calculatorInput.PaymentMethod)
	if err != nil {
		fmt.Println(err)
		result := util.Result(util.ERROR, nil, "类型转换错误！")
		c.Data["json"] = result
		c.ServeJSON()
		return
	}
	switch paymentMethod {
	case 0:
		outPut = service.InterestAvg(entity.CalculatorInput{
			Name:          calculatorInput.Name,
			Investment:    investment,
			Proportion:    proportion,
			Years:         years,
			InterestRate:  interestRate,
			UpProportion:  upProportion,
			PaymentMethod: paymentMethod,
			StartMonth:    calculatorInput.StartMonth,
		})
	case 1:
		outPut = service.MoneyAvg(entity.CalculatorInput{
			Name:          calculatorInput.Name,
			Investment:    investment,
			Proportion:    proportion,
			Years:         years,
			InterestRate:  interestRate,
			UpProportion:  upProportion,
			PaymentMethod: paymentMethod,
			StartMonth:    calculatorInput.StartMonth,
		})
	default:
		c.Data["json"] = util.Result(util.ERROR, nil, "请求解析失败！")
		c.ServeJSON()
		return
	}

	result := util.Result(util.SUCCESS, outPut, "success")
	c.Data["json"] = result
	c.ServeJSON()
}
