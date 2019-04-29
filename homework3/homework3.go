package main

/*
 * Syntax Go Homework 3
 * Stepanov Anton, 29.04.2019
 */

import "fmt"

type Vehicle struct {
	Type              string
	Vendor            string
	Model             string
	Year              int
	VolumeTrunk       int
	VolumeTrunkFilled int
	EngineRunning     bool
	WindowsOpen       bool
	EcologyClass      string
}

var queue []string //очередь

func main() {

	truckMan := Vehicle{
		"truck",
		"MAN",
		"TGX",
		2019,
		30000,
		10,
		false,
		true,
		"Euro5",
	}

	fmt.Println(truckMan)

	truckMan.VolumeTrunkFilled = 29000

	truckMan.EngineRunning = true

	truckMan.WindowsOpen = false

	fmt.Println(truckMan)

	crossoverHyundai := Vehicle{
		"crossover",
		"Hyundai",
		"SantaFe",
		2015,
		585,
		350,
		false,
		false,
		"Euro5",
	}

	fmt.Println(crossoverHyundai)

	crossoverHyundai.VolumeTrunkFilled = 0

	fmt.Println(crossoverHyundai)

	fmt.Println(queue)

	QueuePush("Первый")

	QueuePush("Второй")

	QueuePush("Последний")

	fmt.Println(queue)

	fmt.Println(QueueShift())

	fmt.Println(QueueShift())

	fmt.Println(QueueShift())

	fmt.Println(QueueShift())

}

/*
 * Добавление в очередь
 */
func QueuePush(newElement string) {
	queue = append(queue, newElement)
}

/*
 * Получение первого элемента
 */
func QueueShift() string {
	numsOfElements := len(queue)
	if numsOfElements == 0 {
		return ""
	}
	ret := queue[0]

	newQueue := make([]string, 0)
	for i, v := range queue {
		if i != 0 {
			newQueue = append(newQueue, v)
		}
	}

	queue = newQueue

	return ret
}
