package excel

import "testing"

func TestRead(t *testing.T) {
	ReadXlsx("D:\\test\\ne\\doc\\test.xlsx")
}

func TestReadXls(t *testing.T) {
	ReadXls("D:\\test\\ne\\doc\\苏宁全量地址数据v+1.4.xls")
}
