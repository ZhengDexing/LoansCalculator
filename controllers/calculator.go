package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"

	"LoansCalculator/entity"
	"LoansCalculator/service"
	"LoansCalculator/util"
	"github.com/astaxie/beego"
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
	hour, err := strconv.Atoi(calculatorInput.Hour)
	price, err := strconv.ParseFloat(calculatorInput.Price, 64)
	powerProportion, err := strconv.ParseFloat(calculatorInput.PowerProportion, 64)
	capacity, err := strconv.Atoi(calculatorInput.Capacity)
	if err != nil {
		fmt.Println(err)
		result := util.Result(util.ERROR, nil, "类型转换错误！")
		c.Data["json"] = result
		c.ServeJSON()
		return
	}

	param := entity.CalculatorInput{
		Name:            calculatorInput.Name,
		Investment:      investment,
		Proportion:      proportion,
		Years:           years,
		InterestRate:    interestRate,
		UpProportion:    upProportion,
		PaymentMethod:   paymentMethod,
		StartMonth:      calculatorInput.StartMonth,
		Hour:            hour,
		Price:           price,
		PowerProportion: powerProportion,
		Capacity:        capacity,
	}

	switch paymentMethod {
	case 0:
		outPut = service.InterestAvg(param)
	case 1:
		outPut = service.MoneyAvg(param)
	default:
		c.Data["json"] = util.Result(util.ERROR, nil, "请求解析失败！")
		c.ServeJSON()
		return
	}
	result := util.Result(util.SUCCESS, outPut, "success")
	c.Data["json"] = result
	c.ServeJSON()
}
