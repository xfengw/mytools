package excelTools

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"log"
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

func WriteListToexcelBody(writer *excelize.File, list [][]string, rowStart int,sheetName string ) {


	//first := writer.Sheets[0]
	//cell :='A'
	start := 0
	rowwidth :=0
	high := len(list)
	if len(list)>0{
		rowwidth=len(list[0])
	}


	for i, row := range list {
		start = rowStart + i
		for j, value := range row {
			//把行坐标与竖坐标转换为cell名，如B3
			cellName := GetNumToExcelABCD(j) + strconv.Itoa(start)
			writer.SetCellValue(sheetName, cellName, value)
		}

	}
	SetBodyStyle(writer,sheetName,rowStart,rowwidth,high)
}
func TopSetDate(excel *excelize.File,sheetName string,toplist []string){
	//写入第一行标题
	len:= len(toplist)
	for i,cell := range toplist{
		excel.SetCellValue(sheetName,GetNumToExcelABCD(i)+strconv.Itoa(1),cell)
	}
	Topstyle(excel,sheetName,len)
}

func Topstyle(excel *excelize.File,sheetName string,len int) {
	style, err := excel.NewStyle(`{
           "font": {"bold": true,"size": 14},
           "fill": {"type": "pattern","color": ["#E0EBF5"],"pattern": 1 },
          "alignment":{"horizontal": "center", "vertical": "center","wrap_text":true},

         "border":
               [{"type":"left","color":"000000","style":2},
                 {"type":"top","color":"000000","style":2},
                 {"type":"bottom","color":"000000","style":2},
                 {"type":"right","color":"000000","style":2}
                 ]
          
    }`)
	if err != nil {
		log.Fatal("ERROR", err.Error())
	}
	excel.SetCellStyle(sheetName,"A1",GetNumToExcelABCD(len-1)+strconv.Itoa(1),style)


}
func SetBodyStyle(excel *excelize.File,sheetName string ,rowStart,len,high int) {

	style, err := excel.NewStyle(`{
         "font": {"bold": false,"size": 10},
         "fill": {"type": "pattern","color": ["#E0EBF5"],"pattern": 1 },
         "alignment":{"horizontal": "left", "vertical": "center","wrap_text":true},
         "border":
               [{"type":"left","color":"000000","style":1},
                {"type":"top","color":"000000","style":1},
                {"type":"bottom","color":"000000","style":1},
                {"type":"right","color":"000000","style":1}
                ]
     }`)
	if err != nil {
		log.Fatal("BodyStyle Error!!", err.Error())
	}
	//rows:=excel.GetRows(sheetName)
	if len>0{
		excel.SetCellStyle(sheetName,"A"+strconv.Itoa(rowStart),GetNumToExcelABCD(len-1)+strconv.Itoa(high+rowStart-1),style)
	}
}
