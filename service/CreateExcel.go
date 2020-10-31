package service

import (
	"bytes"
	"strconv"

	"LoansCalculator/entity"
	"LoansCalculator/util"
	"github.com/360EntSecGroup-Skylar/excelize"
)

type ExcelFile struct {
	CalculatorInput entity.CalculatorInput
	Result          entity.Result
}

// excel 输入输出内容写入excel中
func (ef ExcelFile) createInOutPut(excelFile *excelize.File) {
	excelFile.SetCellValue("Sheet1", "A1", "输入内容")
	excelFile.SetCellValue("Sheet1", "A2", "风场名字")
	excelFile.SetCellValue("Sheet1", "B2", ef.CalculatorInput.Name)
	excelFile.SetCellValue("Sheet1", "A3", "风电场总投资(万元)")
	excelFile.SetCellValue("Sheet1", "B3", ef.CalculatorInput.Investment)
	excelFile.SetCellValue("Sheet1", "A4", "自投比例(100%)")
	proportion := strconv.FormatFloat(ef.CalculatorInput.Proportion, 'f', -1, 64) + "%"
	excelFile.SetCellValue("Sheet1", "B4", proportion)
	excelFile.SetCellValue("Sheet1", "A5", "融资年限(年)")
	excelFile.SetCellValue("Sheet1", "B5", ef.CalculatorInput.Years)
	excelFile.SetCellValue("Sheet1", "A6", "融资基本利率(100%)")
	interestRate := strconv.FormatFloat(ef.CalculatorInput.InterestRate, 'f', -1, 64) + "%"
	excelFile.SetCellValue("Sheet1", "B6", interestRate)
	excelFile.SetCellValue("Sheet1", "A7", "融资上浮比例(100%)")
	upProportion := strconv.FormatFloat(ef.CalculatorInput.UpProportion, 'f', -1, 64) + "%"
	excelFile.SetCellValue("Sheet1", "B7", upProportion)
	excelFile.SetCellValue("Sheet1", "A8", "还款方式")
	if ef.CalculatorInput.PaymentMethod == 0 {
		excelFile.SetCellValue("Sheet1", "B8", "等额本息")
	} else if ef.CalculatorInput.PaymentMethod == 1 {
		excelFile.SetCellValue("Sheet1", "B8", "等额本金")
	}
	excelFile.SetCellValue("Sheet1", "A9", "还款起始月份")
	excelFile.SetCellValue("Sheet1", "B9", ef.CalculatorInput.StartMonth)
	excelFile.SetCellValue("Sheet1", "A10", "可研满发小时(h)")
	excelFile.SetCellValue("Sheet1", "B10", ef.CalculatorInput.Hour)
	excelFile.SetCellValue("Sheet1", "A11", "电价(元)")
	excelFile.SetCellValue("Sheet1", "B11", ef.CalculatorInput.Price)
	excelFile.SetCellValue("Sheet1", "A12", "限电比例(100%)")
	powerProportion := strconv.FormatFloat(ef.CalculatorInput.PowerProportion, 'f', -1, 64) + "%"
	excelFile.SetCellValue("Sheet1", "B12", powerProportion)
	excelFile.SetCellValue("Sheet1", "A13", "风场容量(kW)")
	excelFile.SetCellValue("Sheet1", "B13", ef.CalculatorInput.Capacity)

	excelFile.SetCellValue("Sheet1", "A15", "输出内容")
	excelFile.SetCellValue("Sheet1", "A16", "自投金额(万元)")
	excelFile.SetCellValue("Sheet1", "B16", ef.Result.CalculatorOutput.Money)
	excelFile.SetCellValue("Sheet1", "A17", "支付总本金额(万元)")
	excelFile.SetCellValue("Sheet1", "B17", ef.Result.CalculatorOutput.LoanMoney)
	excelFile.SetCellValue("Sheet1", "A18", "融资月利率(100%)")
	interestRateOut := strconv.FormatFloat(ef.Result.CalculatorOutput.InterestRate, 'f', -1, 64) + "%"
	excelFile.SetCellValue("Sheet1", "B18", interestRateOut)
	excelFile.SetCellValue("Sheet1", "A19", "还款总额(万元)")
	excelFile.SetCellValue("Sheet1", "B19", ef.Result.CalculatorOutput.Repayment)
	excelFile.SetCellValue("Sheet1", "A20", "支付总利息(万元)")
	excelFile.SetCellValue("Sheet1", "B20", ef.Result.CalculatorOutput.Interest)
	excelFile.SetCellValue("Sheet1", "A21", "融资月数(月)")
	excelFile.SetCellValue("Sheet1", "B21", ef.Result.CalculatorOutput.Month)
	excelFile.SetCellValue("Sheet1", "A22", "月均还款(万元)")
	excelFile.SetCellValue("Sheet1", "B22", ef.Result.CalculatorOutput.MonthRepayment)
	excelFile.SetCellValue("Sheet1", "A23", "年发电收益(万元)")
	excelFile.SetCellValue("Sheet1", "B23", ef.Result.CalculatorOutput.IncomeYear)
	excelFile.SetCellValue("Sheet1", "A24", "20年发电收益(万元)")
	excelFile.SetCellValue("Sheet1", "B24", ef.Result.CalculatorOutput.Income20)
	excelFile.SetCellValue("Sheet1", "A25", "年收益率(100%)")
	incomeRate := strconv.FormatFloat(ef.Result.CalculatorOutput.IncomeRate, 'f', -1, 64) + "%"
	excelFile.SetCellValue("Sheet1", "B25", incomeRate)
	excelFile.SetCellValue("Sheet1", "A26", "预计回本年限(年)")
	excelFile.SetCellValue("Sheet1", "B26", ef.Result.CalculatorOutput.Year)
}

// 月还款计划
func (ef ExcelFile) createMonthPlant(excelFile *excelize.File) {
	excelFile.SetCellValue("Sheet1", "A28", "月序号")
	excelFile.SetCellValue("Sheet1", "B28", "月份")
	excelFile.SetCellValue("Sheet1", "C28", "归还本金(万元)")
	excelFile.SetCellValue("Sheet1", "D28", "归还利息(万元)")
	excelFile.SetCellValue("Sheet1", "E28", "本息合计(万元)")
	excelFile.SetCellValue("Sheet1", "F28", "剩余本金")
	excelFile.SetCellValue("Sheet1", "G28", "月结余(万元)")

	var line = 29
	for _, value := range ef.Result.RepaymentPlan {
		column := strconv.FormatInt(int64(line), 10)
		excelFile.SetCellValue("Sheet1", "A"+column, value[0])
		excelFile.SetCellValue("Sheet1", "B"+column, value[1])
		excelFile.SetCellValue("Sheet1", "C"+column, value[2])
		excelFile.SetCellValue("Sheet1", "D"+column, value[3])
		excelFile.SetCellValue("Sheet1", "E"+column, value[4])
		excelFile.SetCellValue("Sheet1", "F"+column, value[5])
		excelFile.SetCellValue("Sheet1", "G"+column, value[6])
		line++
	}
}

func (ef ExcelFile) analysis(excelFile *excelize.File) error {
	arr := ef.Result.RepaymentPlan

	excelFile.SetCellValue("Sheet1", "K1", "预计年还本金利息")
	excelFile.SetCellValue("Sheet1", "K2", arr[0][4].(float64)*12)
	excelFile.SetCellValue("Sheet1", "L1", "年发电量收益")
	excelFile.SetCellValue("Sheet1", "L2", ef.Result.CalculatorOutput.IncomeYear)
	excelFile.SetCellValue("Sheet1", "J4", "年")
	excelFile.SetCellValue("Sheet1", "K4", "风场融资累计支出(万元)")
	excelFile.SetCellValue("Sheet1", "L4", "风场发电累计收益(万元)")

	var line = 5
	for i := 1; i <= 20; i++ {
		column := strconv.FormatInt(int64(line), 10)
		excelFile.SetCellValue("Sheet1", "J"+column, i)
		f := util.Round(ef.CalculatorInput.Investment + ef.Result.CalculatorOutput.Interest/
			float64(ef.CalculatorInput.Years)*float64(i))
		excelFile.SetCellValue("Sheet1", "K"+column, f)
		profit := ef.Result.CalculatorOutput.IncomeYear * float64(i)
		excelFile.SetCellValue("Sheet1", "L"+column, profit)
		line++
	}
	err := excelFile.AddChart("Sheet1", "L28", `{"type":"pie","series":[{"categories":"Sheet1!$A$16:$A$17","values":"Sheet1!$B$16:$B$17"}],"format":{"x_scale":1.0,"y_scale":1.0,"x_offset":15,"y_offset":10,"print_obj":true,"lock_aspect_ratio":false,"locked":false},"legend":{"position":"bottom","show_legend_key":false},"title":{"name":"投资情况分析"},"plotarea":{"show_bubble_size":true,"show_cat_name":false,"show_leader_lines":false,"show_percent":false,"show_series_name":false,"show_val":false},"show_blanks_as":"gap"}`)
	if err != nil {
		return err
	}
	err = excelFile.AddChart("Sheet1", "T28", `{"type":"pie","series":[{"categories":"Sheet1!$A$17,Sheet1!$A$20","values":"Sheet1!$B$17,Sheet1!$B$20"}],"format":{"x_scale":1.0,"y_scale":1.0,"x_offset":15,"y_offset":10,"print_obj":true,"lock_aspect_ratio":false,"locked":false},"legend":{"position":"bottom","show_legend_key":false},"title":{"name":"融资成本分析"},"plotarea":{"show_bubble_size":true,"show_cat_name":false,"show_leader_lines":false,"show_percent":false,"show_series_name":false,"show_val":false},"show_blanks_as":"gap"}`)
	if err != nil {
		return err
	}
	//err = excelFile.AddChart("Sheet1", "T28", `{"type":"pie","series":[{"categories":"Sheet1!$A$17,$A$20","values":"Sheet1!$B$17,$B$20"}],"format":{"x_scale":1.0,"y_scale":1.0,"x_offset":15,"y_offset":10,"print_obj":true,"lock_aspect_ratio":false,"locked":false},"legend":{"position":"bottom","show_legend_key":false},"title":{"name":"融资成本分析"},"plotarea":{"show_bubble_size":true,"show_cat_name":false,"show_leader_lines":false,"show_percent":false,"show_series_name":false,"show_val":false},"show_blanks_as":"gap"}`)
	err = excelFile.AddChart("Sheet1", "L44", `{"type":"col","series":[
									{"name":"Sheet1!$K$4","categories":"Sheet1!$J$5:$J$24","values":"Sheet1!$K$5:$K$24"},
									{"name":"Sheet1!$L$4","categories":"Sheet1!$J$5:$J$24","values":"Sheet1!$L$5:$L$24"}],
									"format":{"x_scale":1.0,"y_scale":1.0,"x_offset":15,"y_offset":10,"print_obj":true,"lock_aspect_ratio":false,"locked":false},"legend":{"position":"left","show_legend_key":false},"title":{"name":"累计支出与收入对比(每年)"},"plotarea":{"show_bubble_size":false,"show_cat_name":false,"show_leader_lines":false,"show_percent":false,"show_series_name":false,"show_val":false},"show_blanks_as":"zero"}`)

	if err != nil {
		return err
	}
	err = excelFile.AddChart("Sheet1", "T44", `{"type":"col","series":[{"categories":"Sheet1!$K$1:$L$1","values":"Sheet1!$K$2:$L$2"}],"format":{"x_scale":1.0,"y_scale":1.0,"x_offset":15,"y_offset":10,"print_obj":true,"lock_aspect_ratio":false,"locked":false},"legend":{"position":"left","show_legend_key":false},"title":{"name":"银行年还款与总发电收益对比(万元)"},"plotarea":{"show_bubble_size":true,"show_cat_name":false,"show_leader_lines":false,"show_percent":true,"show_series_name":true,"show_val":false},"show_blanks_as":"zero"}`)
	if err != nil {
		return err
	}
	return nil
}

func (ef ExcelFile) CreateExcelFilm() (*bytes.Buffer, error) {
	var excelFile = excelize.NewFile()
	excelFile.SetColWidth("Sheet1", "A", "G", 18)
	excelFile.SetColWidth("Sheet1", "J", "L", 15)
	style, err := excelFile.NewStyle(`{"alignment":{"horizontal":"right"}}`)
	if err != nil {
		return nil, err
	}
	excelFile.SetCellStyle("Sheet1", "B1", "B26", style)
	ef.createInOutPut(excelFile)
	ef.createMonthPlant(excelFile)
	err = ef.analysis(excelFile)
	buffer, err := excelFile.WriteToBuffer()
	return buffer, err
}
