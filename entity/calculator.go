package entity

// 计算器输入实体
type CalculatorInput struct {
	Name          string  //风电场名称
	Investment    float64 //风电场总投资（万元）
	Proportion    float64 //自投比例
	Years         int     //贷款年限
	InterestRate  float64 //基本利率
	UpProportion  float64 //上浮比例
	PaymentMethod int     //还款方式
	StartMonth    string  //还款起始月份
}

// 计算器输出实体
type CalculatorOutput struct {
	Money          float64 // 自投金额
	LoanMoney      float64 // 贷款金额
	InterestRate   float64 // 贷款月利率 （100%）
	Repayment      float64 // 还款总金额
	Interest       float64 // 支付利息
	Month          int     // 贷款月数
	MonthRepayment float64 // 月均还款
}

// 月还款计划实体
type RepaymentPlan struct {
	Time  string  // 年月
	Money float64 //金额
}

// 成功返回的实体
type Result struct {
	CalculatorOutput CalculatorOutput
	RepaymentPlan    []RepaymentPlan
}
