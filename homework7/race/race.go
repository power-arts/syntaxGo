package main

/*
 * Syntax Go Homework 7
 * Stepanov Anton, 24.05.2019
 * 5 пункт
 */

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var startFlag bool = false
var finish chan int
var raceTrack int = 1000

type Barrier struct {
	NumOfThreads int
	wg           sync.WaitGroup
}

func NewBarrier(num int) (b *Barrier) {
	b = &Barrier{NumOfThreads: num}
	b.wg.Add(num)
	return
}

func (b *Barrier) Await() {
	b.wg.Wait()
	time.Sleep(time.Millisecond)
	b.wg.Add(1)
}

func (b *Barrier) Done() {
	b.wg.Done()
}

func car(num int, barrier *Barrier) {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	var speed int = r.Perm(100)[5] + 10

	fmt.Println("Машина[", num, "] Начало подготовки, speed: ", speed)

	time.Sleep(time.Duration(r.Perm(5)[4]) * time.Second)

	fmt.Println("Машина[", num, "] закончила подготовку")

	barrier.Done()

	barrier.Await()

	for {
		if startFlag {
			break
		}
		time.Sleep(10 * time.Millisecond)
	}

	fmt.Println("Машина[", num, "] приступила к гонке")

	track := raceTrack

	for {
		track -= speed
		if track <= 0 {
			break
		}
		time.Sleep(100 * time.Millisecond)
	}

	fmt.Println("Машина[", num, "] Финишировала")

	barrier.Done()

	finish <- num

}

func main() {
	barrier := NewBarrier(5)
	for i := 0; i < barrier.NumOfThreads; i++ {
		go car(i, barrier)
	}

	barrier.wg.Wait()
	fmt.Println("Прошли подготовку")

	time.Sleep(500 * time.Millisecond)

	for i := 5; i > 0; i-- {
		fmt.Println(i)
		time.Sleep(1000 * time.Millisecond)
	}

	fmt.Println("GO!")

	finish = make(chan int, 4)

	startFlag = true

	barrier.wg.Wait()

	fmt.Println("Результаты гонки:")

	for place := 1; place <= 5; place++ {
		select {
		case car := <-finish:
			fmt.Println(place, " place - Машина: ", car)
		}
	}

	close(finish)

}
