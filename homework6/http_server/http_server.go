package main

/*
 * Syntax Go Homework 6
 * Stepanov Anton, 18.05.2019
 * 3 пункт
 */

import (
	"log"
	"net/http"
	"text/template"
)

func hello(res http.ResponseWriter, req *http.Request) {

	const tmpl = `<doctype html>
	<html>
		<head>
			<title>Hello {{.Name}}!</title>
		</head>
		<body>
			Hello {{.Name}}!
		</body>
	</html>`
	req.ParseForm()
	res.Header().Set("Content-Type", "text/html")

	data := struct {
		Name string
	}{}

	data.Name = req.Form.Get("name")

	t := template.Must(template.New("").Parse(tmpl))

	t.Execute(res, data)
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	http.HandleFunc("/hello", hello)

	err := http.ListenAndServe(":80", nil)
	log.Println("Error creating http server: ", err)
}
