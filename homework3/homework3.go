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

}
