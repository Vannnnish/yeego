/**
 * Created by angelina on 2017/7/29.
 */

package yeeExcel

import "github.com/tealeg/xlsx"

// 传入excel的文件地址
// 第一个参数  是excel表格第一行的数据数组，一般为解释性type
// 第二个参数  是excel中真正的数据
// 第三个参数  为二维数组形式返回是否错误
func ReadExcel(filePath string) ([]string, [][]string, error) {
	firstRow := make([]string, 0)
	otherRows := make([][]string, 0)
	file, err := xlsx.OpenFile(filePath)
	if err != nil {
		return firstRow, otherRows, err
	}
	for _, sheet := range file.Sheets {
		for k, row := range sheet.Rows {
			if k == 0 {
				for _, cell := range row.Cells {
					str, _ := cell.String()
					firstRow = append(firstRow, str)
				}
			} else {
				otherRow := make([]string, 0)
				for _, cell := range row.Cells {
					str, _ := cell.String()
					otherRow = append(otherRow, str)
				}
				otherRows = append(otherRows, otherRow)
			}
		}
	}
	return firstRow, otherRows, nil
}

// 创建excel
// 第一个参数  是excel表格第一行的数据数组，一般为解释性type
// 第二个参数  是excel中真正的数据
// 第三个参数  存储文件的路径
func CreateExcel(firstRow []string, otherRows [][]string, path string) error {
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Sheet1")
	if err != nil {
		return err
	}
	//添加首行数据
	row := sheet.AddRow()
	columns := len(firstRow)
	for i := 0; i < columns; i++ {
		cell := row.AddCell()
		cell.Value = firstRow[i]
	}
	//添加剩余的行数数据
	rows := len(otherRows)
	for i := 0; i < rows; i++ {
		row = sheet.AddRow()
		for j := 0; j < columns; j++ {
			cell := row.AddCell()
			cell.Value = otherRows[i][j]
		}
	}
	err = file.Save(path)
	if err != nil {
		return err
	}
	return nil
}
