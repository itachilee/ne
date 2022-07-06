package excel

import (
	"fmt"

	"github.com/extrame/xls"
	"github.com/xuri/excelize/v2"
)

func ReadXlsx(path string) {
	// f, err := excelize.OpenFile("D:\\test\\ne\\doc\\苏宁全量地址数据v+1.4.xls")
	f, err := excelize.OpenFile(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// Get value from cell by given worksheet name and axis.
	cell, err := f.GetCellValue("Sheet1", "B2")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cell)
	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}
}

func ReadXls(file_path string) (res [][]string) {
	if xlFile, err := xls.Open(file_path, "utf-8"); err == nil {
		fmt.Println(xlFile.Author)
		//第一个sheet
		sheet := xlFile.GetSheet(0)
		if sheet.MaxRow != 0 {
			temp := make([][]string, sheet.MaxRow)
			for i := 0; i < int(sheet.MaxRow); i++ {
				row := sheet.Row(i)
				data := make([]string, 0)
				if row.LastCol() > 0 {
					for j := 0; j < row.LastCol(); j++ {
						col := row.Col(j)
						data = append(data, col)
					}
					temp[i] = data
				}
			}
			res = append(res, temp...)
		}
	} else {
		fmt.Print(err)
	}
	return res
}
