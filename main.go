package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("message.txt")
	if err != nil {
		fmt.Println(err)
	}
	for {
		storage := make([]byte, 8)
		n, err := f.Read(storage)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("read: %s\n", string(storage[:n]))
	}
}
