package main

import (
	"encoding/csv"
	"fmt"
	"sort"
	"strconv"
)

type Processor struct {
	header []string
	grades [][]string
	writer *csv.Writer
}

type grades [][]string

var idx = 0

func (x grades) Len() int {
	return len(x)
}

func (x grades) Less(i, j int) bool {
	a, _ := strconv.Atoi(x[i][idx])
	b, _ := strconv.Atoi(x[j][idx])
	return a < b
}

func (x grades) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

// Handle get data from generator. When dataCh is close, sort data and print
func (p Processor) Handle(dataCh <-chan []string) {
	go func() {
		for {
			grade, ok := <-dataCh
			if ok {
				p.grades = append(p.grades, grade)
			} else {
				p.Print()
				break
			}
		}
	}()
}

func (p Processor) Print() {
	fmt.Println("###### sort by gradeA ######")
	p.Sort(0)
	p.writer.WriteAll(append([][]string{p.header}, p.grades...))
	fmt.Println("###### sort by gradeB ######")
	p.Sort(1)
	p.writer.WriteAll(append([][]string{p.header}, p.grades...))
	fmt.Println("###### sort by gradeC ######")
	p.Sort(2)
	p.writer.WriteAll(append([][]string{p.header}, p.grades...))
}

func (p Processor) Sort(flag int) {
	idx = flag
	sort.Sort(grades(p.grades))
}
