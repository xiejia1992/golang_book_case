package main

import (
	"html/template"
	"os"
	"log"
)

func main()  {
	const templ = `<p>A: {{.A}}</p><p>B: {{.B}}</p>`
	t := template.Must(template.New("escape").Parse(templ))
	var date struct{
		A string
		B template.HTML
	}
	date.A = "<b>Hello!</b>"
	date.B = "<b>Hello!</b>"
	if err := t.Execute(os.Stdout, date); err != nil {
		log.Fatal(err)
	}
}