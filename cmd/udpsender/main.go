package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	UDPAddress, err := net.ResolveUDPAddr("udp", "localhost:42069")
	conn, err := net.DialUDP("udp", nil, nil)
	if err != nil {
		fmt.Println("error while UDP")
	}
	reader := bufio.NewReader(os.Stdin)

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
