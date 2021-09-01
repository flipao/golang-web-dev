package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"text/template"
	"time"
)

type TableRow struct {
	Date     time.Time
	Open     float64
	High     float64
	Low      float64
	Close    float64
	Volume   int
	AdjClose float64
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	data := loadData()

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatal(err)
	}
}

func loadData() []TableRow {
	r, err := readCsv("table.csv")
	if err != nil {
		log.Fatal("Error reading csv", err)
	}

	tr := []TableRow{}
	for _, row := range r[1:] {
		r0, _ := time.Parse("2006-01-02", row[0])
		r1, _ := strconv.ParseFloat(row[1], 64)
		r2, _ := strconv.ParseFloat(row[2], 64)
		r3, _ := strconv.ParseFloat(row[3], 64)
		r4, _ := strconv.ParseFloat(row[4], 64)
		r5, _ := strconv.ParseInt(row[5], 10, 64)
		r6, _ := strconv.ParseFloat(row[6], 32)
		tr = append(tr, TableRow{
			Date:     r0,
			Open:     r1,
			High:     r2,
			Low:      r3,
			Close:    r4,
			Volume:   int(r5),
			AdjClose: r6,
		})
	}
	return tr
}

func readCsv(f string) ([][]string, error) {
	fh, err := os.Open(f)
	if err != nil {
		return [][]string{}, err
	}
	defer fh.Close()

	rows, err := csv.NewReader(fh).ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return rows, nil
}
