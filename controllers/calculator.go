package controllers

import (
	"LoansCalculator/entity"
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
	var calculatorInput entity.CalculatorInput
	var plans []entity.RepaymentPlan
	// 获取requestBody
	data := c.Ctx.Input.RequestBody
	// json解析成对象
	err := json.Unmarshal(data, &calculatorInput)
	if err != nil {
		result := util.Result(util.ERROR, nil, "请求解析失败！")
		c.Data["json"] = result
		c.ServeJSON()
	}

	fmt.Println("calculatorInput:", calculatorInput)

	// 自投金额 = 总投资X自投比例
	money := calculatorInput.Investment * (calculatorInput.Proportion / 100)
	// 贷款金额
	loanMoney, _ := strconv.ParseFloat(
		util.BigNumberSub(calculatorInput.Investment, money),
		64,
	)
	// 贷款月数
	month := calculatorInput.Years * 12
	calculatorOutputRes := entity.CalculatorOutput{
		Money:     money,
		LoanMoney: loanMoney,
		Month:     month,
	}

	monthList := util.GetAllMonth(calculatorInput.StartMonth, month)

	for _, value := range monthList {
		plans = append(plans, entity.RepaymentPlan{Time: value, Money: 123213})
	}

	outPut := entity.Result{
		CalculatorOutput: calculatorOutputRes,
		RepaymentPlan:    plans,
	}

	result := util.Result(util.SUCCESS, outPut, "success")
	c.Data["json"] = result
	c.ServeJSON()
}
