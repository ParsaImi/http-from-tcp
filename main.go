package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("message.txt")
	if err != nil {
		fmt.Println(err)
	}

	line := ""
	storage := make([]byte, 8)
	for {
		n, err := f.Read(storage)
		if err != nil {
			fmt.Println(err)
			break
		}
		data := storage[:n]
		if i := bytes.IndexByte(data, '\n'); i != -1 {
			line += string(storage[:i])
			data = data[i+1:]
			fmt.Printf("read: %s\n", line)
			line = ""
		}
		line += string(data)
	}
}
