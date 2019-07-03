package main

import (
	"fmt"
	"strings"
	U "jason-gao/golangUsing/util"
)

func main() {

	// string
	//https://golang.org/pkg/strings/#LastIndex

	ip := "12.2.111.255"
	fmt.Println(strings.Index("chicken", "ken"))
	fmt.Println(strings.Index("chicken", "dmr"))
	str := "chicken"
	fmt.Println(str[:4])
	fmt.Println("=====Index end======")

	fmt.Println(strings.Index("go gopher", "go"))
	fmt.Println(strings.LastIndex("go gopher", "go"))
	fmt.Println(strings.LastIndex("go gopher", "rodent"))
	pos := strings.LastIndex(ip, ".")
	fmt.Println(pos)
	fmt.Print(ip[:pos])
	fmt.Println("====LastIndex end ====")

	fmt.Println(strings.LastIndexAny("go gopher", "go"))
	fmt.Println(strings.LastIndexAny("go gopher", "rodent"))
	fmt.Println(strings.LastIndexAny("go gopher", "fail"))
	fmt.Println(strings.LastIndexAny(ip, "."))
	pos = strings.LastIndex(ip, ".")
	fmt.Println(ip[:pos])
	fmt.Println("======LastIndexAny end =====")

	//ip
	ipLong, err := U.Ip2long(ip)
	fmt.Println(ipLong, err)

	ip = U.Long2ip(201486335)
	U.P(ip)


}
