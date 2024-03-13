package main

import (
	"fmt"
	"math"
)

type hitung interface {
	luas() float64
	keliling() float64
}

type lingkaran struct {
	diameter float64
}

func (l lingkaran) luas() float64 {
	return math.Pi * math.Pow(l.diameter/2, 2)
}

func (l lingkaran) keliling() float64 {
	return math.Pi * l.diameter
}

func main() {
	var duaDimensi hitung = lingkaran{2.5}
	fmt.Printf("%v %v %v", duaDimensi.keliling(), duaDimensi.luas(), duaDimensi.(lingkaran).diameter)
}
