package entity

// 计算器输入实体
type CalculatorInput struct {
	name          string //风电场名称
	investment    string //风电场总投资（万元）
	proportion    string //自投比例
	years         int    //贷款年限
	interestRate  string //基本利率
	upProportion  string //上浮比例
	paymentMethod string //还款方式
	startMonth    string //还款起始月份
}

// 计算器输出实体
type CalculatorOutput struct {
	money          float64 // 自投金额
	loanMoney      float64 // 贷款金额
	interestRate   float64 // 贷款月利率 （100%）
	repayment      float64 // 还款总金额
	interest       float64 // 支付利息
	month          int     // 贷款月数
	monthRepayment float64 // 月均还款
}

// 月还款计划实体
type repaymentPlan struct {
	time  string  // 年月
	money float64 //金额
}
