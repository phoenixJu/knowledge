package main

import (
	"fmt"
	"net"
)

func main() {
	chat(fmt.Sprintf("Stop %v crawler on %v\n %# v", 3, GetLocalIP(), 6))
}

func chat(message string)  {
	fmt.Printf("hello, %v", message)
}

func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
