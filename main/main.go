package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/kungtalon/boxspacing"
	"github.com/samber/lo"
)

func main() {
	f, err := os.Open("./boxes.csv")
	if err != nil {
		log.Fatal("Unable to read input file ", err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to load as csv ", err)
	}

	rects := lo.Map(records, func(record []string, _ int) boxspacing.Rectangle {
		coords := make([]float64, len(records))
		for idx, col := range record {
			val, _ := strconv.ParseFloat(col, 64)
			coords[idx] = val
		}
		return boxspacing.Rectangle{
			Min: boxspacing.Point{X: coords[0], Y: coords[1]},
			Max: boxspacing.Point{X: coords[2], Y: coords[3]},
		}
	})

	results := boxspacing.Process(rects, nil)
	save(results)
}

func save(rects []boxspacing.Rectangle) {
	f, err := os.Create("./result.csv")
	if err != nil {
		log.Fatal("Unable to read input file ", err)
	}
	defer f.Close()

	csvwriter := csv.NewWriter(f)
	for _, rect := range rects {
		_ = csvwriter.Write([]string{
			fmt.Sprint(rect.Min.X), fmt.Sprint(rect.Min.Y),
			fmt.Sprint(rect.Max.X), fmt.Sprint(rect.Max.Y),
		})
	}
	csvwriter.Flush()
}
