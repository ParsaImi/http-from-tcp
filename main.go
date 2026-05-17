package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("messages.txt")
	if err != nil {
		fmt.Println(err)
	}
	for {
		bytes := make([]byte, 8)
		n, err := file.Read(bytes)
		if err != nil {
			fmt.Println("EOF")
			break
		}
		fmt.Printf("read: %s\n", string(bytes[:n]))

	}

}
