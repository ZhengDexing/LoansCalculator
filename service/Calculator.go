package service

import (
	"LoansCalculator/entity"
	"LoansCalculator/util"
	"fmt"
	"strconv"
)

// 计算等额本息
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
	// 生成excel数据格式
	excel := InterestAvgExcel(calculatorInput.StartMonth, month, loanMoney, interestRate/100)
	calculatorOutput := entity.CalculatorOutput{
		Money:          money,
		LoanMoney:      loanMoney,
		InterestRate:   interestRate,
		Repayment:      repayment,
		Interest:       interest,
		Month:          month,
		MonthRepayment: monthRepayment,
	}

	return entity.Result{
		CalculatorOutput: calculatorOutput,
		RepaymentPlan:    excel,
	}
}

// 等额本金
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

	// 生成excel数据格式
	excel := MoneyAvgExcel(calculatorInput.StartMonth, month, loanMoney, interestRate/100)

	calculatorOutput := entity.CalculatorOutput{
		Money:          money,
		LoanMoney:      loanMoney,
		InterestRate:   interestRate,
		Repayment:      repayment,
		Interest:       interest,
		Month:          month,
		MonthRepayment: monthRepayment,
	}

	return entity.Result{
		CalculatorOutput: calculatorOutput,
		RepaymentPlan:    excel,
	}
}

func InterestAvgExcel(startMonth string, month int, loanMoney float64, interestRate float64) (all []interface{}) {
	fmt.Println(startMonth, "---", month, "---", loanMoney, "---", interestRate)
	months := util.GetAllMonth(startMonth, month)
	var oldLoanMoney = loanMoney
	for key, value := range months {
		fmt.Println(key, "----", value)
		var line []interface{}
		line = append(line, key+1)
		line = append(line, value)
		// 归还本金
		principal := util.Round(loanMoney * interestRate * (util.Powerf2(1+interestRate, key)) /
			(util.Powerf2(1+interestRate, month) - 1))
		line = append(line, principal)
		interest := util.Round(loanMoney * interestRate * (util.Powerf2(1+interestRate, month) - util.Powerf2(1+interestRate, key)) /
			(util.Powerf2(1+interestRate, month) - 1))
		line = append(line, interest)
		sum := util.Round(principal + interest)
		line = append(line, sum)
		principalAll := util.Round(oldLoanMoney - principal)
		line = append(line, principalAll)
		oldLoanMoney = principalAll
		all = append(all, line)
	}
	return all
}

func MoneyAvgExcel(startMonth string, month int, loanMoney float64, interestRate float64) (all []interface{}) {
	months := util.GetAllMonth(startMonth, month)
	var oldSurplus = loanMoney
	for key, value := range months {
		var line []interface{}
		line = append(line, key+1)
		line = append(line, value)
		principal := util.Round(loanMoney / float64(month))
		line = append(line, principal)
		surplus := util.Round(oldSurplus - principal)
		oldSurplus = surplus
		interest := util.Round(surplus * interestRate)
		line = append(line, interest)
		sum := util.Round(principal + interest)
		line = append(line, sum)
		line = append(line, surplus)
		all = append(all, line)
	}
	return all
}
