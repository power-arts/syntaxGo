package main

/*
 * Syntax Go Homework 7
 * Stepanov Anton, 23.05.2019
 * 1 пункт
 */

import (
	"fmt"
	"time"
)

func main() {
	go spinner(50 * time.Millisecond)
	time.Sleep(10 * time.Second)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range "-\\|/" {
			fmt.Printf("%c\r", r)
			time.Sleep(delay)
		}
	}
}
