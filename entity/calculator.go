package entity

// 计算器输入实体
type CalculatorInput struct {
	Name            string  //风电场名称
	Investment      float64 //风电场总投资（万元）
	Proportion      float64 //自投比例
	Years           int     //贷款年限
	InterestRate    float64 //基本利率
	UpProportion    float64 //上浮比例
	PaymentMethod   int     //还款方式
	StartMonth      string  //还款起始月份
	Hour            int     //科研满发小时数（h）
	Price           float64 //电价(元)
	PowerProportion float64 //限电比例（100%）
	Capacity        int     //风场容量(kW)
}

// 计算器输入实体
type CalculatorInputStr struct {
	Name            string //风电场名称
	Investment      string //风电场总投资（万元）
	Proportion      string //自投比例
	Years           string //贷款年限
	InterestRate    string //基本利率
	UpProportion    string //上浮比例
	PaymentMethod   string //还款方式
	StartMonth      string //还款起始月份
	Hour            string //科研满发小时数（h）
	Price           string //电价(元)
	PowerProportion string //限电比例（100%）
	Capacity        string //风场容量(kW)
}

// 计算器输出实体
type CalculatorOutput struct {
	Money          float64     // 自投金额
	LoanMoney      float64     // 贷款金额
	InterestRate   float64     // 贷款月利率 （100%）
	Repayment      float64     // 还款总金额
	Interest       float64     // 支付利息
	Month          int         // 贷款月数
	MonthRepayment float64     // 月均还款
	IncomeYear     float64     // 年发电收益(万元)
	Income20       float64     // 20年发电收益(万元)
	IncomeRate     float64     // 年收益率（100%）
	Year           float64     // 预计回本年限(年)
	PrincipalYear  float64     // 预计年还本金利息
	ProfitYear     float64     // 年发电量收益
	Expenditure    [][]float64 // 支出收入
}

// 月还款计划实体
type RepaymentPlan struct {
	Time  string  // 年月
	Money float64 //金额
}

// 成功返回的实体
type Result struct {
	CalculatorOutput CalculatorOutput
	RepaymentPlan    [][7]interface{}
}
