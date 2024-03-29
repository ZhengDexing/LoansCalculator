package service

import (
	"strconv"

	"LoansCalculator/entity"
	"LoansCalculator/util"
)

// InterestAvg 计算等额本息
func InterestAvg(calculatorInput entity.CalculatorInput) entity.Result {
	// 自投金额 = 总投资X自投比例
	money := util.Round(calculatorInput.Investment * (calculatorInput.Proportion / 100))
	// 贷款金额
	loanMoney, _ := strconv.ParseFloat(
		util.BigNumberSub(calculatorInput.Investment, money),
		64,
	)
	loanMoney = util.Round(loanMoney)
	// 贷款月数
	month := calculatorInput.Years * 12
	interestRate := util.Round(((calculatorInput.InterestRate / 100) * (1 + (calculatorInput.UpProportion / 100)) / 12) * 100)

	// 月均还款
	monthRepayment := util.Round(loanMoney * (interestRate / 100) * util.Powerf2(1+(interestRate/100), month) /
		(util.Powerf2(1+(interestRate/100), month) - 1))
	repayment := util.Round(loanMoney * (interestRate / 100) * util.Powerf2(1+(interestRate/100), month) /
		(util.Powerf2(1+(interestRate/100), month) - 1) * 240)
	interest := util.Round(repayment - loanMoney)
	// 年发电收益
	incomeYear := util.Round(float64(calculatorInput.Capacity) * float64(calculatorInput.Hour) * calculatorInput.Price *
		(1 - (calculatorInput.PowerProportion)/100) / 10000)
	// 20年收益
	income20 := util.Round(incomeYear * 20)
	// 年收益率（100%）
	incomeRate := util.Round((incomeYear / calculatorInput.Investment) * 100)

	//year := int(math.Ceil(repayment / incomeYear))
	year := util.Round(repayment / incomeYear)
	// 生成excel数据格式
	excel := InterestAvgExcel(calculatorInput.StartMonth, month, loanMoney, interestRate/100, incomeYear)
	principal := excel[0][4].(float64) * 12

	var expenditure [][]float64
	for i := 1; i <= 20; i++ {
		var line []float64
		f := util.Round(calculatorInput.Investment + interest/
			float64(calculatorInput.Years)*float64(i))
		profit := incomeYear * float64(i)
		line = append(line, f)
		line = append(line, profit)
		expenditure = append(expenditure, line)
	}
	calculatorOutput := entity.CalculatorOutput{
		Money:          money,
		LoanMoney:      loanMoney,
		InterestRate:   interestRate,
		Repayment:      repayment,
		Interest:       interest,
		Month:          month,
		MonthRepayment: monthRepayment,
		IncomeYear:     incomeYear,
		Income20:       income20,
		IncomeRate:     incomeRate,
		Year:           year,
		PrincipalYear:  principal,
		ProfitYear:     incomeYear,
		Expenditure:    expenditure,
	}

	return entity.Result{
		CalculatorOutput: calculatorOutput,
		RepaymentPlan:    excel,
	}
}

// MoneyAvg 等额本金
func MoneyAvg(calculatorInput entity.CalculatorInput) entity.Result {
	// 自投金额 = 总投资X自投比例
	money := util.Round(calculatorInput.Investment * (calculatorInput.Proportion / 100))
	// 贷款金额
	loanMoney, _ := strconv.ParseFloat(
		util.BigNumberSub(calculatorInput.Investment, money),
		64,
	)
	loanMoney = util.Round(loanMoney)
	// 贷款月利率
	interestRate := util.Round(((calculatorInput.InterestRate / 100) * (1 + (calculatorInput.UpProportion / 100)) / 12) * 100)
	// 贷款月数
	month := calculatorInput.Years * 12

	// 还款总金额
	repayment := ((loanMoney/float64(month) + loanMoney*(interestRate/100)) +
		loanMoney/float64(month)*(1+(interestRate/100))) / 2 * float64(month)
	repayment = util.Round(repayment)
	// 支付利息
	interest := util.Round(repayment - loanMoney)

	monthRepayment := util.Round(repayment / float64(month))

	// 年发电收益
	incomeYear := util.Round(float64(calculatorInput.Capacity) * float64(calculatorInput.Hour) * calculatorInput.Price *
		(1 - (calculatorInput.PowerProportion)/100) / 10000)
	// 20年收益
	income20 := util.Round(incomeYear * 20)
	// 年收益率（100%）
	incomeRate := util.Round((incomeYear / calculatorInput.Investment) * 100)

	//year := int(math.Ceil(repayment / incomeYear))
	year := util.Round(repayment / incomeYear)
	// 生成excel数据格式
	excel := MoneyAvgExcel(calculatorInput.StartMonth, month, loanMoney, interestRate/100, incomeYear)
	principal := excel[0][4].(float64) * 12
	var expenditure [][]float64
	for i := 1; i <= 20; i++ {
		var line []float64
		f := util.Round(calculatorInput.Investment + interest/
			float64(calculatorInput.Years)*float64(i))
		profit := incomeYear * float64(i)
		line = append(line, f)
		line = append(line, profit)
		expenditure = append(expenditure, line)
	}
	calculatorOutput := entity.CalculatorOutput{
		Money:          money,
		LoanMoney:      loanMoney,
		InterestRate:   interestRate,
		Repayment:      repayment,
		Interest:       interest,
		Month:          month,
		MonthRepayment: monthRepayment,
		IncomeYear:     incomeYear,
		Income20:       income20,
		IncomeRate:     incomeRate,
		Year:           year,
		PrincipalYear:  principal,
		ProfitYear:     incomeYear,
		Expenditure:    expenditure,
	}

	return entity.Result{
		CalculatorOutput: calculatorOutput,
		RepaymentPlan:    excel,
	}
}

// InterestAvgExcel 生成计算等额本息 Excel 数据
func InterestAvgExcel(startMonth string, month int, loanMoney float64, interestRate float64, incomeYear float64) (all [][7]interface{}) {
	months := util.GetAllMonth(startMonth, month)
	var oldLoanMoney = loanMoney
	for key, value := range months {
		var line [7]interface{}
		line[0] = key + 1
		line[1] = value
		// 归还本金
		principal := util.Round(loanMoney * interestRate * (util.Powerf2(1+interestRate, key)) /
			(util.Powerf2(1+interestRate, month) - 1))
		line[2] = principal
		interest := util.Round(loanMoney * interestRate * (util.Powerf2(1+interestRate, month) - util.Powerf2(1+interestRate, key)) /
			(util.Powerf2(1+interestRate, month) - 1))
		line[3] = interest
		sum := util.Round(principal + interest)
		line[4] = sum
		principalAll := util.Round(oldLoanMoney - principal)
		if principalAll < 0 {
			line[2] = oldLoanMoney
			line[5] = 0
		} else {
			line[5] = principalAll
		}
		line[6] = util.Round(incomeYear/12 - principal - interest)
		oldLoanMoney = principalAll
		all = append(all, line)
	}
	return all
}

// MoneyAvgExcel 等额本金 Excel 数据
func MoneyAvgExcel(startMonth string, month int, loanMoney float64, interestRate float64, incomeYear float64) (all [][7]interface{}) {
	months := util.GetAllMonth(startMonth, month)
	var oldSurplus = loanMoney
	for key, value := range months {
		var line [7]interface{}
		line[0] = key + 1
		line[1] = value
		principal := util.Round(loanMoney / float64(month))
		line[2] = principal
		surplus := util.Round(oldSurplus - principal)

		interest := util.Round(surplus * interestRate)
		line[3] = interest
		sum := util.Round(principal + interest)
		line[4] = sum

		if surplus < 0 {
			line[2] = oldSurplus
			line[4] = oldSurplus
			line[5] = 0
		} else {
			line[5] = surplus
		}
		line[6] = util.Round(incomeYear/12 - principal - interest)
		oldSurplus = surplus
		all = append(all, line)
	}
	return all
}
