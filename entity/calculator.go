package entity

//计算器输入实体
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
