package util

import (
	"encoding/binary"
	"errors"
	"fmt"
	"net"
)

func Ip2long(ipAddr string) (uint32, error) {
	ip := net.ParseIP(ipAddr)
	if ip == nil {
		return 0, errors.New("wrong ipAddr format")
	}
	ip = ip.To4()
	return binary.BigEndian.Uint32(ip), nil
}

func Long2ip(ipLong uint32) string {
	ipByte := make([]byte, 4)
	binary.BigEndian.PutUint32(ipByte, ipLong)
	ip := net.IP(ipByte)
	return ip.String()
}

func IntToIP(ip uint32) string {
	result := make(net.IP, 4)
	result[0] = byte(ip)
	result[1] = byte(ip >> 8)
	result[2] = byte(ip >> 16)
	result[3] = byte(ip >> 24)
	return result.String()
}

func P(param ...interface{})  {
	fmt.Println(param)
}


func RemoveDuplicates(elements []int) []int {
	// Use map to record duplicates as we find them.
	encountered := map[int]bool{}
	result := make([]int, 0)

	for k := range elements {

		fmt.Println("k", k)
		if encountered[elements[k]] == true {
			// Do not add duplicate.
		} else {
			// Record this element as an encountered element.
			encountered[elements[k]] = true
			// Append to result slice.
			result = append(result, elements[k])
		}
	}
	// Return the new slice.
	return result
}

