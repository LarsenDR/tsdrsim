package main

import (
	"fmt"
	"math/rand"
	"net"
	"strings"
	"time"
)

type Ports struct {
	// provisioning channel
	porta int
	portb int
	// configuration channel
	portc int
	portd int
	//data channel
	porte int
	portf int
}

func randPort(pts Ports) int {
	rand.Seed(time.Now().UnixNano())
	min := 49151
	max := 65535
	// Test number is already been assigned
	value := rand.Intn(max-min+1) + min
	return value
}

func main() {
	var ports Ports
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
		if strings.HasPrefix(data, "TA") { //Discovery
			porta := strings.Split(remote.String(), ":")[1]
			portb := randPort(ports)

			str := fmt.Sprintf("  TangerineSDR PortA %s PortB %d -- From:%s To:%s\n", porta, portb, remote, conn.LocalAddr().String())
			fmt.Printf("%v\n", str)
			reply := []byte(str)
			_, err := conn.WriteToUDP(reply, remote)
			if err != nil {
				panic(err)
			}
		} else if strings.HasPrefix(data, "S?") { //Status
			str := fmt.Sprintf("  TangerineSDR S?\n")
			fmt.Printf("%v\n", str)
			reply := []byte(str)
			_, err := conn.WriteToUDP(reply, remote)
			if err != nil {
				panic(err)
			}
		} else if strings.HasPrefix(data, "T?") { //Time
			t := time.Now()
			str := fmt.Sprintf("  TangerineSDR Time: %v\n", t)
			fmt.Printf("%v\n", str)
			reply := []byte(str)
			_, err := conn.WriteToUDP(reply, remote)
			if err != nil {
				panic(err)
			}
		} else if strings.HasPrefix(data, "C?") { //Channel List
			str := fmt.Sprintf("  TangerineSDR C?\n")
			fmt.Printf("%v\n", str)
			reply := []byte(str)
			_, err := conn.WriteToUDP(reply, remote)
			if err != nil {
				panic(err)
			}
		} else if strings.HasPrefix(data, "P?") { //Channel List
			str := fmt.Sprintf("  TangerineSDR P?\n")
			fmt.Printf("%v\n", str)
			reply := []byte(str)
			_, err := conn.WriteToUDP(reply, remote)
			if err != nil {
				panic(err)
			}
		}

	}
}
