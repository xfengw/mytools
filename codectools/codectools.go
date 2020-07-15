package codectools

import (
	"github.com/axgle/mahonia"
	"encoding/base64"
	"golang.org/x/net/idna"
	"math/rand"
	"strings"
	"time"
)

/*****************************************************************************************
/   字符串按照指定的字符集转码成byte
************************************************************************************** */
func StrToByteByCode(src string, srcCode string, targetCode string) []byte {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(targetCode)
	_, tagByte, _ := tagCoder.Translate([]byte(srcResult), true)
	return tagByte
}
/*****************************************************************************************
/   生成指定长度的包含英文大小写和数字的字符串
************************************************************************************** */
func GenerateRandomString(num int) (randomString string) {
	const randomStringTemp = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890"
	//randomString := ""
	rand.Seed(int64(time.Now().Nanosecond()))
	for i := 0; i < num; i++ {
		i := rand.Intn(62)
		randomString += string(randomStringTemp[i])

	}
	//fmt.Println(randomString,len(randomString))
	return
}
/*****************************************************************************************
/   生成指定长度的包含英文大小写和数字的字符串
************************************************************************************** */


func PunycodeToUnicode(domain string) (newdomain string) {
	newdomain = domain
	var err error
	if strings.Contains(domain, "--") {

		newdomain, err = idna.ToUnicode(domain)
		if err != nil {
			return newdomain
		}
	}
	return newdomain
}
func CodeConvertByEncode(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}

func Base64GbkToUtf8(s string)string{
	if !strings.Contains(s,"=?GB2312?B?"){
		return s
	}
	basestrings := strings.Split(s,"=?GB2312?B?")
	//fmt.Println(basestrings[1])
	basestr :=strings.Split(basestrings[1],"==?=")
	//fmt.Println(basestr[0])
	gbkbyte, _ := base64.RawURLEncoding.DecodeString(basestr[0])
	return CodeConvertByEncode(string(gbkbyte),"gbk","utf-8")
}