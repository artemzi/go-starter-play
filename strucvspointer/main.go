package main

import (
	"log"
	"reflect"
)

func main() {
	type fisk struct {
		Name string
	}

	f1 := fisk{Name: "kat"}
	v := reflect.ValueOf(&f1)
	y := (*v.Interface().(*fisk))

	log.Print(reflect.DeepEqual(f1, y))
}