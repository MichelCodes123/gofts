package gofts

import (
	"fmt"
	"reflect"
	"testing"
)

type str struct {
	A int32
	B string
	C []int
	D []string
}

func TestFTS(t *testing.T) {

	var s str
	var g str
	data := make(map[string][]string)
	data["A"] = []string{"1"}
	data["B"] = []string{"Bonjour"}
	data["C"] = []string{"1", "2", "3"}
	data["D"] = []string{"Peter", "James", "John"}

	g.A = 1
	g.B = "Bonjour"
	g.C = []int{1, 2, 3}
	g.D = []string{"Peter", "James", "John"}

	e := Fts(data, &s)
	fmt.Print(e)

	r := reflect.DeepEqual(s, g)

	if !r{
		t.Fatal("Struct is not persisted properly")
	}

}
