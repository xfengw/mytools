package excelTools

import (
	"fmt"
	"testing"
)

func TestGetNumToExcelABCD(t *testing.T) {
	for num :=0;num<1000;num++{
		fmt.Println(num,GetNumToExcelABCD(num))
	}

}
func TestExcelTimeToUnixTime(t *testing.T) {
	exceltime :=43906
	fmt.Println(ExcelTimeToUnixTime(exceltime).Format("2006-01-02"))
}