package main

import (
	"log"
	"net"
)

/*
 * Syntax Go Homework 7
 * Stepanov Anton, 24.05.2019
 * 4 пункт
 */

func main() {
	reactHost := mirroredQuery()
	log.Println(reactHost)
}

func mirroredQuery() string {
	responses := make(chan string, 3)
	go func() {
		responses <- request("yandex.ru")
	}()
	go func() {
		responses <- request("google.ru")
	}()
	go func() {
		responses <- request("nic.xyz")
	}()
	return <-responses // возврат самого быстрого ответа
}

func request(hostname string) string {
	conn, err := net.Dial("tcp", hostname+":80")
	if err != nil {
		log.Println(hostname, " unreachable")
	}
	defer conn.Close()
	return hostname
}
