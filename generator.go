package main

import (
	"k8s.io/apimachinery/pkg/util/wait"
	"math/rand"
	"strconv"
	"time"
)

type Generator struct {
}

func (g Generator) Generate(dataCh chan<- []string, stopCh <-chan struct{}) {
	rand.Seed(time.Now().UnixNano())
	wait.Until(func() {
		dataCh <- []string{
			strconv.Itoa(rand.Intn(100)),
			strconv.Itoa(rand.Intn(100)),
			strconv.Itoa(rand.Intn(100)),
		}
	}, time.Second, stopCh)
}
