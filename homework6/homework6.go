package main

/*
 * Syntax Go Homework 6
 * Stepanov Anton, 18.05.2019
 * 2,4,5 пункты
 */

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math"
	"os"
)

var teal color.Color = color.RGBA{0, 200, 200, 255}
var red color.Color = color.RGBA{200, 30, 30, 255}

var chessBorderColor color.Color = color.RGBA{204, 204, 204, 255}

func main() {

	//Создать изображение с вертикальными и горизонтальными линиями
	createImage()

	//Создать шахматную доску
	createChessBoard()

	fmt.Println("Решение квадратных уравнений")

	var a float64
	var b float64
	var c float64

	fmt.Println("Введите a:")
	fmt.Scanln(&a)

	fmt.Println("Введите b:")
	fmt.Scanln(&b)

	fmt.Println("Введите c:")
	fmt.Scanln(&c)

	x1, x2, err := quadraticEquations(a, b, c)
	fmt.Println("X1=", x1, " X2=", x2, " Error:", err)
}

//Пункт 2 ДЗ
func createImage() {

	img := image.NewRGBA(image.Rect(0, 0, 500, 500))

	draw.Draw(img, img.Bounds(), &image.Uniform{teal}, image.ZP, draw.Src)

	//Горизонтальные линии
	for x := 50; x <= 450; x++ {
		y := 35
		img.Set(x, y, red)
		img.Set(x, y+3, red)
		img.Set(x, 500-y, red)
		img.Set(x, 500-y+3, red)
	}

	//Вертикальные линии
	for y := 50; y <= 450; y++ {
		x := 35
		img.Set(x, y, red)
		img.Set(x+3, y, red)
		img.Set(500-x, y, red)
		img.Set(500-x+3, y, red)
	}

	//Создаем файл
	file, err := os.Create("someimage.png")
	if err != nil {
		fmt.Errorf("%s", err)
	}
	defer file.Close()

	//Пишем в файл
	png.Encode(file, img)

	fmt.Println("Line image created.")

}

//Пункт 5 ДЗ
func createChessBoard() {

	boardImage := image.NewRGBA(image.Rect(0, 0, 900, 900))

	draw.Draw(boardImage, boardImage.Bounds(), &image.Uniform{chessBorderColor}, image.ZP, draw.Src)

	squareMask := image.NewRGBA(image.Rect(0, 0, 100, 100))

	draw.Draw(squareMask, squareMask.Bounds(), image.White, image.ZP, draw.Src)

	var xLineCord int = 50
	var yLineCord int = 50

	var lastColor image.Image = image.Black

	for i := 1; i <= 64; i++ {

		draw.DrawMask(boardImage, image.Rect(xLineCord, yLineCord, xLineCord+100, yLineCord+100), lastColor, image.ZP, squareMask, image.ZP, draw.Over)

		lastColor = changeColor(lastColor)

		xLineCord += 100

		if i%8 == 0 {
			//переход на новую линию y
			yLineCord += 100
			//возврат к началу по x
			xLineCord = 50
			//внеочередная смена цвета клетки
			lastColor = changeColor(lastColor)
		}

	}

	//Создаем файл
	file, err := os.Create("chessboard.png")
	if err != nil {
		fmt.Errorf("%s", err)
	}
	defer file.Close()

	//Пишем в файл
	png.Encode(file, boardImage)

	fmt.Println("Chessboard created.")

}

func changeColor(lastColor image.Image) image.Image {
	if lastColor == image.White {
		return image.Black
	}
	return image.White
}

func quadraticEquations(a float64, b float64, c float64) (x1 float64, x2 float64, err error) {

	discriminant := b*b - 4*a*c

	if discriminant > 0 {
		x1 = (-b + math.Sqrt(discriminant)) / (2 * a)
		x2 = (-b - math.Sqrt(discriminant)) / (2 * a)
	} else if discriminant == 0 {
		x1 = (-b + math.Sqrt(discriminant)) / (2 * a)
	} else if discriminant < 0 {
		err = errors.New("Корней нет")
	}

	return

}
