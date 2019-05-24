package main

/*
 * Syntax Go Homework 7
 * Stepanov Anton, 23.05.2019
 * 2 пункт
 */

import (
	"fmt"
	"time"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// генерация
	go func() {
		for x := 0; x < 10; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	// возведение в квадрат
	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	// печать
	for x := range squares {
		fmt.Println(x)
		time.Sleep(100 * time.Millisecond)
	}
}
