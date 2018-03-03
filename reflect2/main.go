package main

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
)

type Name struct {
	First, Last string
}

func (n *Name) String() string {
	return n.First + " " + n.Last
}

func main() {
	// === reflect ===
	n := &Name{First: "Indigo", Last: "Montoya"}
	stringer := (*fmt.Stringer)(nil) // pointer to nil of type Stringer
	implements(n, stringer)

	writer := (*io.Writer)(nil)
	implements(n, writer)
	// ===

	// === type===
	b := bytes.NewBuffer([]byte("Hello"))
	if isStringer(b) {
		fmt.Printf("(typed) %T is a stringer\n", b)
	}

	i := 123
	if isStringer(i) {
		fmt.Printf("(typed) %T is a stringer\n", i)
	}
	// ===
}

func implements(concrete interface{}, target interface{}) bool {
	iface := reflect.TypeOf(target).Elem()

	v := reflect.ValueOf(concrete)
	t := v.Type() // reflect.Type

	if t.Implements(iface) {
		fmt.Printf("(reflect) %T is a %s\n", concrete, iface.Name())
		return true
	}
	fmt.Printf("(reflect) %T is not a %s\n", concrete, iface.Name())
	return false
}

// check interface based on type
func isStringer(v interface{}) bool {
	_, ok := v.(fmt.Stringer)
	return ok
}
