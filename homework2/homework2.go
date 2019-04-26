package main

import "fmt"

func main() {
	var action string

	fmt.Println("Выберите действие: (even, modulo, fibo, primes)")
	fmt.Scanln(&action)

	switch action {
	default:
		fmt.Printf("Неизвестная команда: %v\nДоступные действия: even, modulo, fibo, primes\n", action)
	case "even":
		{
			var number int
			fmt.Println("Введите целое число:")
			fmt.Scanln(&number)
			if modulo(number, 2) {
				fmt.Printf("Число %v чётное\n", number)
			} else {
				fmt.Printf("Число %v нечётное\n", number)
			}
		}
	case "modulo":
		{
			var number int
			var divider = 3

			fmt.Println("Введите целое число:")
			fmt.Scanln(&number)

			if modulo(number, divider) {
				fmt.Printf("Число %v делится на %v\n", number, divider)
			} else {
				fmt.Printf("Число %v не детится на %v\n", number, divider)
			}
		}
	case "fibo":
		{
			var answer int
			fmt.Println("Вывести на экран или заполнить срез? 1|2")
			fmt.Scanln(&answer)
			switch answer {
			case 1:
				fibo1()
			case 2:
				fmt.Println(fibo2())
			default:
				{
					fmt.Println("Ответ не разпознан, будет выведен на экран")
					fibo1()
				}
			}

		}
	case "primes":
		{
			var number int

			fmt.Println("Введите целое число:")
			fmt.Scanln(&number)

			fmt.Println(createSlice(number))
		}
	}
}

func modulo(number, divider int) bool {
	if number%divider == 0 {
		return true
	}
	return false
}

func fibo1() {
	var number int
	fmt.Println("Функция которая последовательно выводит на экран N первых чисел Фибоначчи, начиная от 0")
	fmt.Println("Введите N:")
	fmt.Scanln(&number)

	var num1 int = 0
	var num2 int = 1

	fmt.Printf("[%v, ", num1)
	fmt.Printf("%v, ", num2)

	for i := 3; i <= number; i++ {
		nowNum := num1 + num2
		if i == number {
			fmt.Printf("%v]\n", nowNum)
		} else {
			fmt.Printf("%v, ", nowNum)
		}
		num1 = num2
		num2 = nowNum
	}

}

func fibo2() []int {
	var number int
	fmt.Println("Функция которая заполняет срез числами Фибоначчи")
	fmt.Println("Введите N:")
	fmt.Scanln(&number)

	arrFibo := make([]int, number)

	arrFibo[0] = 0
	arrFibo[1] = 1

	fmt.Println(arrFibo)

	for i := 2; i < number; i++ {
		arrFibo[i] = arrFibo[i-1] + arrFibo[i-2]
	}

	return arrFibo
}

func isPrime(number int) bool {
	if number == 1 {
		return false
	}

	for d := 2; d*d <= number; d++ {
		if number%d == 0 {
			return false
		}
	}

	return true
}

func createSlice(n int) []int {
	arrPrime := make([]int, n)
	var number int = 1
	var counter int = 0

	for counter < n {
		if isPrime(number) {
			arrPrime[counter] = number
			counter++
		}
		number++
	}

	return arrPrime
}
