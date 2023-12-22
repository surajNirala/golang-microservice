package main

import (
	"github.com/surajNirala/go-customers/app"
)

type sumInterface interface {
	sum() int
}

type sumStruct struct {
	first  int
	second int
}

func (v sumStruct) sum() int {
	return v.first + v.second
}

func main() {
	app.Start()
}
