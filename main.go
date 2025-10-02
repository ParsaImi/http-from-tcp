package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("messages.txt")
	if err != nil {
		fmt.Println("error while reading file")
	}
	currentline := ""
	for {
		data := make([]byte, 8)
		n, err := file.Read(data)
		if err != nil {
			break
		}
		data = data[:n]                                // dont get this one
		if i := bytes.IndexByte(data, '\n'); i != -1 { // dont get this one
			currentline += string(data[:i])
			data = data[i+1:]
			fmt.Printf("read: %s\n", currentline)
			currentline = ""
		}
		currentline += string(data)
	}
	if len(currentline) != 0 {
		fmt.Printf("read: %s\n", currentline)
	}

}
