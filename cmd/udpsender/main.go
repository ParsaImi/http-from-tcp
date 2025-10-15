package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
)

func main() {
	UDPAddress, err := net.ResolveUDPAddr("udp", "localhost:42069")
	conn, err := net.DialUDP("udp", UDPAddress, UDPAddress)
	if err != nil {
		fmt.Println("error while UDP")
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("error while tcp")
		}
		for line := range getLinesChannel(conn) {
			fmt.Printf("read: %s\n", line)
		}
	}

}
