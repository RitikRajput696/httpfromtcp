package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
)

func getLinesChannel(file io.ReadCloser) <-chan string {
	out := make(chan string, 1)

	go func() {
		defer close(out)
		defer file.Close()
		str := ""

		for {
			data := make([]byte, 8)
			n, err := file.Read(data)

			if err != nil {
				break
			}
			data = data[:n]
			if i := bytes.IndexByte(data, '\n'); i != -1 { // we found a '\n' char ( newline)
				str += string(data[:i])
				data = data[i+1:]

				out <- str
				str = ""
			}

			str += string(data)

		}

		if len(str) != 0 {

			out <- str
		}

	}()

	return out
}

func main() {

	ln, err := net.Listen("tcp", ":42069")
	if err != nil {
		fmt.Println("failed to listen", err)
		return
	}

	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("failed to Accept connection", err)
			return
		}

		// fmt.Println("connection has been established")
		for line := range getLinesChannel(conn) {
			fmt.Println(line)
		}
	}

}
