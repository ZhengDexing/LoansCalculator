package service

import (
	"LoansCalculator/entity"
	"LoansCalculator/util"
	"strconv"
)

// 计算等额本息
func InterestAvg(calculatorInput entity.CalculatorInput) entity.CalculatorOutput {
	// 自投金额 = 总投资X自投比例
	money := calculatorInput.Investment * (calculatorInput.Proportion / 100)
	// 贷款金额
	loanMoney, _ := strconv.ParseFloat(
		util.BigNumberSub(calculatorInput.Investment, money),
		64,
	)
	// 贷款月数
	month := calculatorInput.Years * 12
	interestRate := ((calculatorInput.InterestRate / 100) * (1 + (calculatorInput.UpProportion / 100)) / 12) * 100

	// 月均还款
	//monthRepayment := loanMoney*(interestRate/100)*util.Powerf2(1+(interestRate/100), month)/
	//	util.Powerf2(1+(interestRate/100), month) - 1
	monthRepayment := loanMoney * (interestRate / 100) * util.Powerf2(1+(interestRate/100), month) /
		(util.Powerf2(1+(interestRate/100), month) - 1)
	repayment := monthRepayment * float64(month)
	interest := repayment - loanMoney
	return entity.CalculatorOutput{
		Money:          money,
		LoanMoney:      loanMoney,
		InterestRate:   interestRate,
		Repayment:      repayment,
		Interest:       interest,
		Month:          month,
		MonthRepayment: monthRepayment,
	}
}

// 等额本金
func MoneyAvg(calculatorInput entity.CalculatorInput) entity.CalculatorOutput {
	// 自投金额 = 总投资X自投比例
	money := calculatorInput.Investment * (calculatorInput.Proportion / 100)
	// 贷款金额
	loanMoney, _ := strconv.ParseFloat(
		util.BigNumberSub(calculatorInput.Investment, money),
		64,
	)
	// 贷款月利率
	interestRate := ((calculatorInput.InterestRate / 100) * (1 + (calculatorInput.UpProportion / 100)) / 12) * 100
	// 贷款月数
	month := calculatorInput.Years * 12

	// 还款总金额
	repayment := ((loanMoney/float64(month) + loanMoney*(interestRate/100)) +
		loanMoney/float64(month)*(1+(interestRate/100))) / 2 * float64(month)

	// 支付利息
	interest := repayment - loanMoney

	monthRepayment := repayment / float64(month)
	return entity.CalculatorOutput{
		Money:          money,
		LoanMoney:      loanMoney,
		InterestRate:   interestRate,
		Repayment:      repayment,
		Interest:       interest,
		Month:          month,
		MonthRepayment: monthRepayment,
	}
}
