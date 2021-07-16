package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		Port: 1024,
		IP:   net.ParseIP("192.168.1.11"),
	})
	if err != nil {
		panic(err)
	}

	defer conn.Close()
	fmt.Printf("server listening %s\n", conn.LocalAddr().String())

	for {
		message := make([]byte, 64)
		rlen, remote, err := conn.ReadFromUDP(message[:])
		if err != nil {
			panic(err)
		}

		
		data := strings.TrimSpace(string(message[:rlen]))
		if strings.HasPrefix(data, "TA") {
			fmt.Printf("%s -- From:%s To:%s\n", data, remote, conn.LocalAddr().String())
		}

	}
}
