package main

import "fmt"
import "math"

const usdRate float64 = 64.05

func main() {
	var action string

	fmt.Println("Выберите действие: (convert, square, deposit)")
	fmt.Scanln(&action)

	switch action {
	default:
		fmt.Printf("Неизвестная команда: %v\nДоступные действия: convert, square, deposit\n", action)
	case "convert":
		convert()
	case "square":
		square()
	case "deposit":
		deposit()
	}
}

func convert() {
	var sumRub float64

	fmt.Printf("Курс доллара: %v\n", usdRate)

	fmt.Println("Введите сумму в рублях:")
	fmt.Scanln(&sumRub)

	fmt.Printf("Сумма в долларах: %.2f\n", sumRub/usdRate)
}

func square() {
	var a float64
	var b float64
	var c float64

	fmt.Println("Введите катет a:")
	fmt.Scanln(&a)

	fmt.Println("Введите катет b:")
	fmt.Scanln(&b)

	c = math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))

	fmt.Printf("Гипотенуза треугольника равна: %.2f\n", c)
	fmt.Printf("Периметр треугольника равна: %.2f\n", a+b+c)
	fmt.Printf("Площадь треугольника равна: %.2f\n", a*b*0.5)
}

func deposit() {
	var depSum float64
	var countYears int
	var capitalization bool
	var capitalizationString string
	var percent float64

	fmt.Println("Введите сумму депозита:")
	fmt.Scanln(&depSum)

	fmt.Println("Введите срок депозита(лет):")
	fmt.Scanln(&countYears)

	fmt.Println("Введите процент по депозиту:")
	fmt.Scanln(&percent)

	fmt.Println("Капитализация (y|n): default: n")
	fmt.Scanln(&capitalizationString)

	if capitalizationString == "y" {
		capitalization = true //Ежемесячная капитализация
	} else {
		capitalization = false
	}

	if capitalization {

		var countMonths int
		countMonths = countYears * 12

		for i := 1; i <= countMonths; i++ {
			depSum += depSum * percent / 100 / 12
			//fmt.Printf("Сумма депозита равна: %.2f\n", depSum)
		}

		fmt.Printf("Финальная сумма депозита равна: %.2f\n", depSum)

	} else {

		for i := 1; i <= countYears; i++ {
			depSum += depSum * percent / 100
			//fmt.Printf("Сумма депозита равна: %.2f\n", depSum)
		}

		fmt.Printf("Финальная сумма депозита равна: %.2f\n", depSum)

	}

}
