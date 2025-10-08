package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
)

func getLinesChannel(f io.ReadCloser) <-chan string {
	out := make(chan string, 1)
	go func() {
		defer f.Close()
		defer close(out)
		currentline := ""
		for {
			data := make([]byte, 8)
			n, err := f.Read(data)
			if err != nil {
				break
			}
			data = data[:n]
			if i := bytes.IndexByte(data, '\n'); i != -1 {
				currentline += string(data[:i])
				data = data[i+1:]
				out <- currentline
				currentline = ""
			}
			currentline += string(data)
		}
		if len(currentline) != 0 {
			out <- currentline
		}
	}()

	return out
}

//	func getLinesChannel(l net.Listener) <-chan string {
//		go func(){
//			defer l.Close()
//			for {
//				l.
//			}
//		}
//	}
func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:42069")
	if err != nil {
		fmt.Println("error while tcp")
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
