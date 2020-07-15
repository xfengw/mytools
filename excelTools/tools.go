package excelTools

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"strconv"
	"time"
)

func GetNumToExcelABCD(num int) string {
	//base := ""
	if num <26{
		return string('A'+num)
	}
	t := num / 26
	m:=num%26

		return GetNumToExcelABCD(t-1)+string('A'+m)

}

func ExcelTimeToUnixTime(exceltime int)time.Time{
	time1:=time.Date(1970,1,3,0,0,0,0,time.Local)
	time2:=time.Date(1900,1,1,0,0,0,0,time.Local)
	duraTime :=time1.Sub(time2)
	baseOriginSecond := int64(duraTime.Microseconds())/1000000
	return time.Unix(int64(exceltime*24*3600)-baseOriginSecond,0)

}

//type excel excelize.File
func WriteStringListToexcel(writer *excelize.File,list [][]string,rowStart int){
	//first := writer.Sheets[0]
	sheetName := "Sheet1"
	//cell :='A'
	start:=0
	for i, row := range list {
		start = rowStart +i
		//fmt.Println()
		for j,value :=range row{
			cellName :=GetNumToExcelABCD(j)+strconv.Itoa(start)
				writer.SetCellValue(sheetName,cellName,value)
		}
	}
}