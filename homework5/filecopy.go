package main

/*
 * Syntax Go Homework 5
 * Stepanov Anton, 12.05.2019
 * 4 пункт
 */

import (
	"flag"
	"io"
	"log"
	"os"
)

func showLog(f bool, text string, num int64) {
	if f {
		if num > 0 {
			log.Printf(text, num)
		} else {
			log.Println(text)
		}
	}
}

func main() {

	sourceFlag := flag.String("s", "", "Source file")
	destinationFlag := flag.String("d", "", "Destination file")
	verboseFlag := flag.Bool("v", false, "Show detail information")
	forceCopyFlag := flag.Bool("f", false, "Copy if destination file exists")

	flag.Parse()

	//check if destination file exists
	if *forceCopyFlag == false {
		if _, err := os.Stat(*destinationFlag); err == nil {
			log.Fatal("Destination file exists")
		}
	}

	showLog(*verboseFlag, "Copy from: "+*sourceFlag+" to: "+*destinationFlag, 0)

	fileSource, err := os.Open(*sourceFlag)
	if err != nil {
		log.Fatal(err)
	}
	defer fileSource.Close()

	fileDestination, err := os.Create(*destinationFlag)
	if err != nil {
		log.Fatal(err)
	}
	defer fileDestination.Close()

	bytesCopied, err := io.Copy(fileDestination, fileSource)
	if err != nil {
		log.Fatal(err)
	}

	showLog(*verboseFlag, "Copied %d bytes.", bytesCopied)

}
