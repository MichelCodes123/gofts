package gofts

import(
	"testing"
	"fmt"
)

	type str struct {
		A int32
		B string
		C []int
		D []string
	}
func TestFTS(t *testing.T){

		var s str
		data := make(map[string][]string)
		data["A"] = []string{"1"}
		data["B"] = []string{"Bonjour"}
		data["C"] = []string{"1", "2", "3"}
		data["D"] = []string{"Peter", "James", "John"}

		e := Fts(data, &s)
		fmt.Println(e)
		fmt.Println(s)

}
