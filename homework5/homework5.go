package main

/*
 * Syntax Go Homework 5
 * Stepanov Anton, 12.05.2019
 * 1,3 пункты задания
 */

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

/*
 * 1) На мой взгляд из статьи "Time in Go: A primer" можно было бы добавить пример про парсинг дат
 */

/*
 * 2) В программе "fileread.go" представленной в методичке чтение производится всего файла целиком,
 * на примере больших файлов это не так плохо если нужно считать весь файл и сохранить его в программе.
 * Но функция Read возвращает длинну прочитанных байт в int и если файл больше 2гб то вернет не верную
 * длинну хотя прочитает файл полностью. Поэтому я сделал чтение файла по частям в цикле до достижения
 * конца файла с проверкой прочитанных байт и длинной файла. А так для обработки больших файлов лучше
 * читать файл по строкам:
 *	scanner := bufio.NewScanner(file)
 *	for scanner.Scan() {
 *		fmt.Println("NEWLINE: ", scanner.Text())
 *	}
 */

type NoteEntity struct {
	Name        string
	Description string
	Time        string
}

func (c NoteEntity) ToSlice() []string {
	row := make([]string, 3, 3)
	row[0] = c.Name
	row[1] = c.Description
	row[2] = c.Time
	return row
}

func main() {

	var notes []NoteEntity

	file, err := os.Open("notes.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	for {
		record, err := reader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		notes = append(notes, NoteEntity{
			record[0],
			record[1],
			record[2],
		})
	}

	fmt.Println(notes)

	for idx, note := range notes {
		notes[idx].Name = note.Name + " NEW"
	}

	fmt.Println(notes)

	sfile, err := os.Create("result.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	w := csv.NewWriter(sfile)

	for _, record := range notes {
		values := record.ToSlice()
		if err := w.Write(values); err != nil {
			log.Fatal(err)
		}
	}

	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}

}
