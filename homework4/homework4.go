package main

/*
 * Syntax Go Homework 4
 * Stepanov Anton, 06.05.2019
 * 1,2,3 пункты задания
 */

import (
	"fmt"
	"sort"

	"./calculator"
	"./phonebook"
)

type Drawable interface {
	Draw()
}

type Pen struct {
	color string
}

func (p Pen) Draw() {
	fmt.Printf("Pen draw, color: %v\n", p.color)
}

type Pencil struct {
	color string
}

func (p Pencil) Draw() {
	fmt.Printf("Pencil draw, color: %v\n", p.color)
}

type Marker struct {
	color string
}

func (p Marker) Draw() {
	fmt.Printf("Marker draw, color: %v\n", p.color)
}

func drawing(drawer Drawable) {
	drawer.Draw()
}

func main() {

	redPen := Pen{"red"}
	drawing(redPen)

	bluePencil := Pencil{"blue"}
	drawing(bluePencil)

	blueMarker := Marker{"green"}
	drawing(blueMarker)

	adressbook := []phonebook.AddressBookEntity{}

	adressbook = append(adressbook, phonebook.AddressBookEntity{
		phonebook.Contact{
			"Anton",
			"Stepanov",
		},
		phonebook.Phones{
			"79020001101",
			"79020001102",
			"79020001103",
			"79020001104",
		},
		"as@ifdev.ru",
		"",
		"",
	})

	adressbook = append(adressbook, phonebook.AddressBookEntity{
		phonebook.Contact{
			"Ivan",
			"Antonov",
		},
		phonebook.Phones{
			"79040001101",
			"79040001102",
			"79040001103",
			"79040001104",
		},
		"ia@ifdev.ru",
		"",
		"",
	})

	adressbook = append(adressbook, phonebook.AddressBookEntity{
		phonebook.Contact{
			"Dmitriy",
			"Sidorov",
		},
		phonebook.Phones{
			"79030001101",
			"79030001102",
			"79030001103",
			"79030001104",
		},
		"sdn@ifdev.ru",
		"",
		"",
	})

	fmt.Println(adressbook)

	//sort by contact.lastname
	sort.Sort(phonebook.AddressBook(adressbook))

	fmt.Println(adressbook)

	fmt.Println("Start calculator:")

	input := ""
	for {
		fmt.Print("> ")
		if _, err := fmt.Scanln(&input); err != nil {
			fmt.Println(err)
			continue
		}

		if input == "exit" {
			break
		}

		if input == "help" {
			calculator.ShowHelp()
			continue
		}

		if res, err := calculator.Calculate(input); err == nil {
			fmt.Printf("Результат: %v\n", res)
		} else {
			fmt.Println("Не удалось произвести вычисление")
		}
	}
}
