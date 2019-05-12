package main

/*
 * Syntax Go Homework 5
 * Stepanov Anton, 12.05.2019
 * 2 пункт задания
 */

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("fileread.go")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// getting size of file
	stat, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}

	fileLength := stat.Size()

	buf := make([]byte, 64*1024*100)

	var fileContent string

	for {
		n, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		fileLength -= int64(n)
		fileContent += string(buf[:n])
	}

	if fileLength != 0 {
		log.Fatal("Read length error")
	}

	fmt.Println(fileContent)

}
