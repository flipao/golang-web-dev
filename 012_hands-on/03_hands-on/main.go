package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type hotel struct {
	Name    string
	Address string
	City    string
	Zip     string
}

type region struct {
	Name   string
	Hotels []hotel
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	c := []region{
		{
			Name: "Southern",
			Hotels: []hotel{
				{"Hotel 1", "10th, Southern ave.", "S. City 1", "123456"},
				{"Hotel 2", "10th, Southern ave.", "S. City 2", "123457"},
			},
		}, {
			Name: "Central",
			Hotels: []hotel{
				{"Hotel 1", "10th, Central ave.", "C. City 1", "234567"},
				{"Hotel 2", "10th, Central ave.", "C. City 2", "234568"},
			},
		}, {
			Name: "Northern",
			Hotels: []hotel{
				{"Hotel 1", "10th, Northern ave.", "N. City 1", "345678"},
				{"Hotel 2", "10th, Northern ave.", "N. City 2", "345679"},
			},
		},
	}

	err := tpl.Execute(os.Stdout, c)
	if err != nil {
		log.Fatal(err)
	}
}
