package ipTools

import (
	"fmt"
	"net"
	"testing"
)

func TestIpIncrement(t *testing.T) {
    fmt.Println(IpIncrement(net.ParseIP("254.255.255.255")))
    //fmt.Print("%b",255)
}


func TestIpDecrement(t *testing.T) {
	fmt.Println(IpDecrement(net.ParseIP("127.1.1.1")))

	//fmt.Print("%b",255)
}

func TestIpCompare(t *testing.T) {
	ip1 :=net.ParseIP("::0:ffff:7f00:01")
	ip2 :=net.ParseIP("127.0.0.1").To4()

	fmt.Println("ip2地址长度：",len(ip2))
	//for i:=0;i<16;i++{
	//	fmt.Println(i,ip1[i],ip2[i])
	//}
	fmt.Println(IpCompare(ip1,ip2))
}

func TestGenerateIpsFromIpSeg(t *testing.T) {
	ip1:=net.ParseIP("50.1.135.3")
	ip2:=net.ParseIP("50.18.115.7")
	ips :=GenerateIpsFromIpSeg(ip1,ip2)
	//fmt.Println("*********",ips)
	for _,ip:=range ips{
		fmt.Println(ip)
	}

}

func TestGenerateIpsFromIpMask(t *testing.T) {

	fmt.Println(GenerateIpsFromIpMask("10.0.0.5/16"))

	//iP,iPnet,_ :=net.I
	//fmt.Println(iP,*iPnet)
}

func TestIpIsInIpSegBymask(t *testing.T) {
	ip :="10.1.255.218"
	ips :="10.0.0.5/16"
	fmt.Println(IpIsInIpSegBymask(ip,ips))
}
