package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
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
	file, err := os.Open("messages.txt")
	if err != nil {
		log.Fatal(err)
	}
	for line := range getLinesChannel(file) {
		fmt.Println("read:", line)
	}

	defer file.Close()

}
