package main

import (
	"encoding/csv"
	"os"
	"time"
)

// main start a generator and processor.
// the generator randomly generates grade and sends to the processor.
// the processor get grade, sort them by different columns and print.
func main() {
	g := Generator{}
	p := Processor{
		header: []string{"gradeA", "gradeB", "gradeC"},
		grades: [][]string{},
		writer: csv.NewWriter(os.Stdout),
	}
	dataCh := make(chan []string)
	stopCh := make(chan struct{})
	go g.Generate(dataCh, stopCh)
	go p.Handle(dataCh)
	time.Sleep(time.Second * 10)
	close(stopCh)
	close(dataCh)
	time.Sleep(time.Second * 5)
}
