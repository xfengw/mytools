package ipTools

import (
	"fmt"
	"log"

	//"encoding/binary"
	"net"
	"strconv"
	"strings"
)
//判断IP地址是否在以-标记的IP段数组中，比如判断Ip是否在电信机房的所有IP段内[112.67.253.192-112.67.253.199 124.225.167.0-124.225.167.255]
func IpIsInIpsegsByline(ip net.IP,ipsegs []string)bool{
    is :=false
    if len(ipsegs)==0{
        return false
    }
    for _,ipseg :=range ipsegs{

        if ipseg==""{continue}
fmt.Println(ip.String(),ipseg)
        is =IpIsInIpsegByline(ip,ipseg)
        if is{return is}
    }
    return is
}
//判断IP地址是否在已掩码长度标记的IP段数组中，比如判断Ip是否在电信机房的所有IP段内[61.186.0.0/18  202.100.192.0/18  218.77.128.0/17]
func IpIsInIpsegsByMask(ip net.IP,ipsegs []string)bool{
    is :=false
    if len(ipsegs)==0{
        return false
    }
    for _,ipseg :=range ipsegs{

        if ipseg==""{continue}
        is =IpNetIsInIpSegBymask(ip,ipseg)
        if is{return is}
    }
    return is
}
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
//fmt.Println("**",ip1.String(),ip2.String())
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
func IpIsInIpsegByStartEndString(ip,ipstart,ipend string)bool{
    return IpIsInIpsegNet(net.ParseIP(ip),net.ParseIP(ipstart),net.ParseIP(ipend))
}
func IpIsInIpSegBymask(ip,ipanmask string)bool{
    ipstart,ipend :=GetNetStartAndEndIp(ipanmask)

    return IpIsInIpsegNet(net.ParseIP(ip),ipstart,ipend)
}

func IpNetIsInIpSegBymask(ip net.IP,ipanmask string)bool{
    ipstart,ipend :=GetNetStartAndEndIp(ipanmask)

    return IpIsInIpsegNet(ip,ipstart,ipend)
}
//判断IP是否在192.168.1.2-192.168.1.254这样的IP段内
func IpIsInIpsegByline(ip net.IP,ips string)bool{
    ips =strings.ReplaceAll(ips," ","")
    ips=strings.Trim(ips,"\r\n")
    if !strings.Contains(ips,"-"){
        if IpCompare(ip,net.ParseIP(ips))==0{
            return true
        }else {
            return false
        }
    }
    ipseg :=strings.Split(ips,"-")
    if len(ipseg)!=2{return false}
    return IpIsInIpsegNet(ip,net.ParseIP(ipseg[0]),net.ParseIP(ipseg[1]))
}
