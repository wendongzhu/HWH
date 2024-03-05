package tool

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"github.com/wendongzhu/hwh/backend/service/global"
	"github.com/wendongzhu/hwh/backend/service/param"
	"os"
	"time"
)

func ReadExcelFile(selectParam param.SelectDataParam) (res param.ExportListParam) {
	fileName := selectParam.FileName
	pageSize := selectParam.PageSize
	currentPage := selectParam.CurrentPage

	filePath := global.ModelFile.ExportFilePath + "/" + fileName
	_, err := os.Stat(filePath)
	fmt.Println(filePath, err)
	if !os.IsNotExist(err) {
		file, err := xlsx.OpenFile(filePath)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			return
		}

		for _, sheet := range file.Sheets {
			for index, row := range sheet.Rows {
				if index > 1 && index > pageSize*(currentPage-1) && index <= pageSize*currentPage {
					var item param.ExportParam
					item.ModelGroup = row.Cells[0].String()
					item.ModelName = row.Cells[1].String()
					item.ProgramNumber = row.Cells[2].String()
					item.WeldTime = row.Cells[3].String()
					item.SamplingTime, _ = row.Cells[4].Int()
					item.ExtractCurrent, _ = row.Cells[5].Int()
					item.ExtractVoltage, _ = row.Cells[6].Int()
					item.ExtractResistance, _ = row.Cells[7].Int()
					item.PeakCurrent, _ = row.Cells[8].Int()
					item.PeakVoltage, _ = row.Cells[9].Int()
					item.PeakResistance, _ = row.Cells[10].Int()
					item.ArchiveSeries = row.Cells[11].String()
					res.Data = append(res.Data, item)
				}
				res.Count = index
			}

		}

	}
	return res
}

func WriteExcelFile(data param.ExportParam) {
	timeStr := time.Now().Format("20060102")
	filePath := global.ModelFile.ExportFilePath + "/" + global.ModelFile.ExportFileName + "_" + timeStr + ".xlsx"

	var file *xlsx.File
	var sheet *xlsx.Sheet
	var err error
	exists, err := PathExists(global.ModelFile.ExportFilePath)
	if err != nil {
		return
	}
	if exists {
		// 检查文件是否存在
		_, err = os.Stat(filePath)
		if os.IsNotExist(err) {
			// 如果文件不存在，创建新的Excel文件
			file = xlsx.NewFile()
			sheet, err = file.AddSheet("Sheet1")
			if err != nil {
				fmt.Printf("Error: %s\n", err)
				return
			}
			row := sheet.AddRow()
			row.AddCell().SetString("ModelGroup")
			row.AddCell().SetString("ModelName")
			row.AddCell().SetString("ProgramNumber")
			row.AddCell().SetString("WeldTime")
			row.AddCell().SetString("SamplingTime")
			row.AddCell().SetString("ExtractCurrent(A)")
			row.AddCell().SetString("ExtractVoltage(V)")
			row.AddCell().SetString("ExtractResistance(μΩ)")
			row.AddCell().SetString("PeakCurrent(A)")
			row.AddCell().SetString("PeakVoltage(V)")
			row.AddCell().SetString("PeakResistance(μΩ)")
			row.AddCell().SetString("ArchiveSeries")
			row = sheet.AddRow()
			row.AddCell().SetString("模块组")
			row.AddCell().SetString("模块名")
			row.AddCell().SetString("程序编号")
			row.AddCell().SetString("焊接时间")
			row.AddCell().SetString("采样时间点")
			row.AddCell().SetString("提取电流(A)")
			row.AddCell().SetString("提取电压(V)")
			row.AddCell().SetString("提取电阻(μΩ)")
			row.AddCell().SetString("电流峰值(A)")
			row.AddCell().SetString("电压峰值(V)")
			row.AddCell().SetString("电阻峰值(μΩ)")
			row.AddCell().SetString("归档序号")

		} else {
			// 如果文件存在，打开该文件
			file, err = xlsx.OpenFile(filePath)
			if err != nil {
				fmt.Printf("Error: %s\n", err)
				return
			}
			// 获取第一个工作表
			sheet = file.Sheets[0]
		}

		row := sheet.AddRow()
		row.AddCell().SetString(data.ModelGroup)
		row.AddCell().SetString(data.ModelName)
		row.AddCell().SetString(data.ProgramNumber)
		row.AddCell().SetString(data.WeldTime)
		row.AddCell().SetInt(data.SamplingTime)
		row.AddCell().SetInt(data.ExtractCurrent)
		row.AddCell().SetInt(data.ExtractVoltage)
		row.AddCell().SetInt(data.ExtractResistance)
		row.AddCell().SetInt(data.PeakCurrent)
		row.AddCell().SetInt(data.PeakVoltage)
		row.AddCell().SetInt(data.PeakResistance)
		row.AddCell().SetString(data.ArchiveSeries)

		err = file.Save(filePath)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		}
	}

}
