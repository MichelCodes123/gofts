package gofts

import (
	"fmt"
	"reflect"
	"testing"
)

// Different Struct types for testing
type str struct {
	A int32
	B string
	C []int
	D []string
}

func testdata() map[string][]string {

	data := make(map[string][]string)
	data["A"] = []string{"1"}
	data["B"] = []string{"Bonjour"}
	data["C"] = []string{"1", "2", "3"}
	data["D"] = []string{"Peter", "James", "John"}

	return data

}

func TestFTS(t *testing.T) {

	var s str
	var g str
	data := testdata()

	g.A = 1
	g.B = "Bonjour"
	g.C = []int{1, 2, 3}
	g.D = []string{"Peter", "James", "John"}

	e := Fts(data, &s)
	fmt.Print(e)

	r := reflect.DeepEqual(s, g)
	if !r {
		t.Fatal("Struct is not persisted properly")
	}
}

func TestInvalidInput(t *testing.T) {
	var s str
	e := Fts(testdata(), s)

	if e == nil {
		t.Fatal("Function allows non pointer to struct")
	}
}

type testIgnore struct {
	A int
	B int8
	C str
	D []str
}

func TestIgnoredType(t *testing.T) {
	var s testIgnore
	data := make(map[string][]string)
	data["A"] = []string{"1"}
	data["B"] = []string{"25"}
	data["C"] = []string{"This is the voice"}
	data["D"] = []string{"God is good"}

	correct := testIgnore{1, 25, str{}, nil}

	e := Fts(data, &s)

	if e != nil {
		t.Fatal(e)
	}
	r := reflect.DeepEqual(s, correct)
	if !r {
		fmt.Println(s)
		fmt.Println(correct)
		t.Fatal("Form data persisted incorrectly.")
	}
}

type allSupported struct {
	A  string
	B  float32
	C  float64
	D  int
	E  int8
	F  int16
	G  int32
	H  int64
	I  uint
	J  uint8
	K  uint16
	L  uint32
	M  uint64
	N  bool
	As []string
	Bs []float32
	Cs []float64
	Ds []int
	Es []int8
	Fs []int16
	Gs []int32
	Hs []int64
	Is []uint
	Js []uint8
	Ks []uint16
	Ls []uint32
	Ms []uint64
	Ns []bool
}

func TestAllSupportedTypes(t *testing.T) {
	var s allSupported
	data := make(map[string][]string)
	data["A"] = []string{"God is good, All the Time"}
	data["B"] = []string{"25.32"}
	data["C"] = []string{"256.322"}
	data["D"] = []string{"1"}
	data["E"] = []string{"25"}
	data["F"] = []string{"256"}
	data["G"] = []string{"2560"}
	data["H"] = []string{"25600"}
	data["I"] = []string{"256000"}
	data["J"] = []string{"23"}
	data["K"] = []string{"25"}
	data["L"] = []string{"256"}
	data["M"] = []string{"2560"}
	data["N"] = []string{"true"}
	data["As"] = []string{"1", "2", "3"}
	data["Bs"] = []string{"23.5", "23.5", "24.5"}
	data["Cs"] = []string{"2342.299"}
	data["Ds"] = []string{"1", "2", "3"}
	data["Es"] = []string{"24", "25", "26"}
	data["Fs"] = []string{"256", "257"}
	data["Gs"] = []string{"2567", "25000"}
	data["Hs"] = []string{"26090", "17000"}
	data["Is"] = []string{"1", "2", "3"}
	data["Js"] = []string{"1", "2", "3"}
	data["Ks"] = []string{"1", "2", "34", "5"}
	data["Ls"] = []string{"256", "257"}
	data["Ms"] = []string{"190234"}
	data["Ns"] = []string{"true", "false"}

	as := []string{"1", "2", "3"}
	bs := []float32{23.5, 23.5, 24.5}
	cs := []float64{2342.299}
	ds := []int{1, 2, 3}
	es := []int8{24, 25, 26}
	fs := []int16{256, 257}
	gs := []int32{2567, 25000}
	hs := []int64{26090, 17000}
	is := []uint{1, 2, 3}
	js := []uint8{1, 2, 3}
	ks := []uint16{1, 2, 34, 5}
	ls := []uint32{256, 257}
	ms := []uint64{190234}
	ns := []bool{true, false}

	correct := allSupported{"God is good, All the Time", 25.32, 256.322, 1, 25, 256, 2560, 25600, 256000, 23, 25, 256, 2560, true, as, bs, cs, ds, es, fs, gs, hs, is, js, ks, ls, ms, ns}

	e := Fts(data, &s)
	if e != nil {
		fmt.Print(e)
	}
	r := reflect.DeepEqual(s, correct)
	if !r {
		fmt.Println(s)
		fmt.Println("Perfect")
		fmt.Println(correct)
		t.Fatal("Form data persisted incorrectly.")
	}

}

type ms1 struct {
	A string
	B int
	C []int
}
type ms2 struct {
	D ms1
	E float32
	F bool
}
type ms3 struct {
	G int
	H float32
	I float64
}

func TestMFTS(t *testing.T) {
	var a ms1
	var b ms2
	var c ms3

	data := make(map[string][]string)

	data["A"] = []string{"This is the voice"}
	data["B"] = []string{"25"}
	data["C"] = []string{"1", "2", "3"}
	data["D"] = []string{"GG's"}
	data["E"] = []string{"25.32", "223.45"}
	data["F"] = []string{"true","false"}
	data["G"] = []string{"1"}
	data["H"] = []string{"25.32"}
	data["I"] = []string{"256782.72"}

	correct1 := ms1{"This is the voice", 25, []int{1, 2, 3}}
	correct2 := ms2{ms1{}, 25.32, true}
	correct3 := ms3{1, 25.32, 256782.72}

	e := Mfts(data, &a, &b, &c)

	if e != nil {
		t.Fatal(e)
	}

	r1 := reflect.DeepEqual(a, correct1)
	r2 := reflect.DeepEqual(b, correct2)
	r3 := reflect.DeepEqual(c, correct3)

	if !r1 || !r2 || !r3 {
		fmt.Println(a)
		fmt.Println(correct1)

		fmt.Println(b)
		fmt.Println(correct2)

		fmt.Println(c)
		fmt.Println(correct3)
		t.Fatal("Form data persisted incorrectly.")
	}

}
