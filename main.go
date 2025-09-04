// package main

// import (
// 	"bytes"
// 	"fmt"
// 	"io"
// 	"log"
// 	"os"
// )

// func main() {

// 	file, err := os.Open("messages.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer file.Close()

// 	strs := ""

// 	for {
// 		dataChunk := make([]byte, 8)
// 		n, err := file.Read(dataChunk)
// 		if err != nil {
// 			break
// 		}

// 		dataChunk = dataChunk[:n]

// 		if i := bytes.IndexByte(dataChunk, '\n'); i != -1 { // we are at the end of line
// 			strs += string(dataChunk[:i])
// 			dataChunk = dataChunk[i+1:]

// 			fmt.Printf("read: %s\n", string(strs))
// 			strs = ""
// 		}

// 		strs += string(dataChunk)

// 	}
// 	if len(strs) != 0 {
// 		fmt.Printf("read: %s\n", strs)
// 	}

// }

// // func getLinesChannel(f io.ReadCloser) <-chan string {

// // }

package main

import (
	"fmt"
)

func main() {
	messages := make(chan string)

	go func() {
		messages <- "ping"
	}()

	msg := <-messages
	fmt.Println(msg)

}
