package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":42069")
	if err != nil {
		fmt.Println(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
		}
		lines := getLinesFromReader(conn)
		for line := range lines {
			fmt.Println(line)
		}
	}
}

func getLinesFromReader(f io.ReadCloser) <-chan string {
	ch := make(chan string, 1)
	go func() {
		defer f.Close()
		defer close(ch)
		line := ""
		storage := make([]byte, 8)
		for {
			n, err := f.Read(storage)

			if n > 0 {
				data := storage[:n]
				// Process all newlines in the chunk
				for i := bytes.IndexByte(data, '\n'); i != -1; i = bytes.IndexByte(data, '\n') { // this is a magic for loop fr
					line += string(data[:i])
					ch <- line
					line = ""
					data = data[i+1:]
				}
				// Append remaining part of the chunk to the line
				line += string(data)
			}

			if err != nil {
				if err == io.EOF {
					break // End of file
				}
				fmt.Println(err) // Other errors
				break
			}
		}
		// Send the last line if it's not empty
		if len(line) > 0 {
			ch <- line
		}
	}()
	return ch

}
