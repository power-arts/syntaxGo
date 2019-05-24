package main

/*
 * Syntax Go Homework 7
 * Stepanov Anton, 24.05.2019
 * 3 пункт
 */

import (
	"io"
	"log"
	"net"
	"os"
	"sync"
	"sync/atomic"
	"time"
)

var connections int32 = 0
var isStop int = 0
var cancel chan int
var wg sync.WaitGroup

func main() {
	cancel = make(chan int)

	wg.Add(1)
	go listen()

	for {
		byte := make([]byte, 4)
		os.Stdin.Read(byte)
		time.Sleep(1 * time.Second)
		if string(byte) == "exit" {
			log.Println("Exiting...")
			if connections > 0 {
				cancel <- 1
			}
			isStop = 1
			wg.Done()
			break
		} else if string(byte) == "conn" {
			log.Println("Active connections: ", connections)
		}
	}

	log.Println("Waiting close connections...")
	wg.Wait()
	log.Println("Application closed...")
}

func listen() {
	defer wg.Done()
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		if isStop > 0 {
			log.Println("Application whil be closed...")
			return
		}
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
		wg.Add(1)
		atomic.AddInt32(&connections, 1)
		log.Print("New connection, active connections: ", connections)
	}
}

func handleConn(c net.Conn) {
	defer func() {
		c.Close()
		atomic.StoreInt32(&connections, connections-1)
		wg.Done()
	}()
	for {
		select {
		case <-cancel:
			isStop = 1
		default:
			_, err := io.WriteString(c, time.Now().Format("15:04:05\n\r"))
			if err != nil {
				return
			}
			time.Sleep(1 * time.Second)
			if isStop > 0 {
				log.Print("Close connection: ", c.RemoteAddr())
				return
			}
		}
	}
}
