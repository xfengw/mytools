package iptools

import (
	"fmt"
	"log"

	//"encoding/binary"
	"net"
	"strconv"
	"strings"
)
//IP地址递增
func IpIncrement(ip net.IP) net.IP{
	len :=len(ip)
	for i:=len-1;i>=0;i--{
		byteTemp :=(ip[i])
	    if byteTemp!=255{
	    	ip[i]=(byte)(byteTemp+1)
	    	return ip
		}else if ip[0]==(byte)(255){
			return net.IPv6zero
		}else{
			ip[i]=(byte)(0)
		}
	}
	return ip
}

//IP地址递减
func IpDecrement(ip net.IP) net.IP{
	len :=len(ip)
	for i:=len-1;i>=0;i--{
		byteTemp :=(ip[i])
		//判断i循环到0时，对应的字节是否为0
		if i==0&&byteTemp==0{
			return net.IPv6zero
		}
		if byteTemp!=0{
			ip[i]=(byte)(byteTemp-1)
			return ip
		}else {
			ip[i]=(byte)(255)
		}
	}
	return ip
}
//IP地址比较
func IpCompare(ip1,ip2 net.IP) int{
   ip1=ip1.To16()
   ip2=ip2.To16()
   for i:=0;i<net.IPv6len;i++{
   	  if ip1[i]!=ip2[i]{
   	  	if ip1[i]<ip2[i]{return -1}else{return 1}
	  }
	}
	return 0
}
//根据开始、结束IP生成所有IP
func GenerateIpsFromIpSeg(start,end net.IP)[]string {
	start =start.To16()
	end=end.To16()
	if IpCompare(start,end)>0{
		start,end = end,start
	}
	//fmt.Println(start,end)
	ips :=make([]string,0,1024)
	for startIp:=start;IpCompare(startIp,end)<=0;startIp=IpIncrement(startIp){
		ips=append(ips,startIp.String())
	}
	return  ips
}
//根据子网掩码位数获取子网起始和终止IP所有IP(未实现
func GetNetStartAndEndIp(ipAndMask string)(startIp,endIp net.IP) {
	ipAndMask = strings.ReplaceAll(ipAndMask," ","")
	ipmask :=strings.Split(ipAndMask,"/")
	ip :=ipmask[0]
	startIp =net.ParseIP(ip)
	endIp =net.ParseIP(ip)
	mask,err :=strconv.Atoi(ipmask[1])
	if err!=nil ||startIp==nil||mask>128{
		log.Println(" iptools.GenerateIpsFromIpMask（）输入参数错误")
		return nil,nil
	}
	unmask :=0
	if mask<32{
		unmask =32-mask
	}else{
		unmask =128-mask
	}
	startIp =startIp.To16()
	for i:=len(startIp)-1;i>=0;i--{
		if unmask <8{
			startIp[i]=(startIp[i]>>unmask)<<unmask
			endIp[i] = endIp[i]|(1<<(unmask)-1)
			break
		}else{
			//fmt.Println(i,iP[i])
			startIp[i] =byte(0)
			endIp[i] =byte(255)
			unmask=unmask-8
		}
	}
	return
}
//根据IP和子网掩码位数生成所有IP(未实现
func GenerateIpsFromIpMask(ipAndMask string)[]string {
	startIp,endIp :=GetNetStartAndEndIp(ipAndMask)
	if startIp==nil||endIp==nil{
		fmt.Println("iptools.GenerateIpsFromIpMask输入参数错误！")
		return []string{}
	}
	return GenerateIpsFromIpSeg(startIp,endIp)

}

func IpIsInIpsegNet(ip,ipstart,ipend net.IP)bool{
	if IpCompare(ipstart,ipend)==1{
		ipstart,ipend=ipend,ipstart
	}
	if IpCompare(ip,ipstart)==-1{return false	}
	if IpCompare(ip,ipend)==1{return false}
	return true
}
func IpIsInIpsegString(ip,ipstart,ipend string)bool{
   return IpIsInIpsegNet(net.ParseIP(ip),net.ParseIP(ipstart),net.ParseIP(ipend))
}
func IpIsInIpSegBymask(ip,ipanmask string)bool{
	ipstart,ipend :=GetNetStartAndEndIp(ipanmask)

	return IpIsInIpsegNet(net.ParseIP(ip),ipstart,ipend)
}
