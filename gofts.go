package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func Hello() string {
	return "hi i've just imported some packages"

}

// Transfers form data into the struct specified by the user. Returns false if the operation fails, and error message for invalid inputs
// Throws error if:
// The dest is not a pointer to a struct
func Fts(form map[string][]string, dest interface{}) error {

	//Verify if the destination string is a struct
	v := reflect.ValueOf(dest)
	if v.Kind() != reflect.Ptr {
		return fmt.Errorf("%s", "Must be a pointer to a struct")
	}

	d := v.Elem()
	if d.Kind() != reflect.Struct {
		return fmt.Errorf("%s", "Not value of struct")
	}

	for i := 0; i < d.NumField(); i++ {
		//f := d.Field(i)
		//Find the name of the field, within the struct
		name := d.Type().Field(i).Name

		//Set the value of the struct field, to the corresponding value with the map
		g, o := form[name]

		if o {
			//Form data mappings may contain a list of values. If the mapping has a length greater than 0.
			//Make sure that the field
			if len(g) > 0 {
				if d.Type().Field(i).Type.Kind() == reflect.Array {
					d.Field(i).Set(reflect.ValueOf(g))
				}
			}
			//Must perform type conversion if the struct value is not of type string.
			converted := tc(d.Field(i).Kind(), g[0])
			d.Field(i).Set(reflect.ValueOf(converted).Convert(d.Field(i).Type()))
		}
	}

	fmt.Print(dest)
	return nil
}
func tc(a reflect.Kind, str string) interface{} {
	g := a.String()
	if a == reflect.String {
		return str
	}
	switch a {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		gg, _ := strconv.Atoi(g[3:])
		i, _ := strconv.ParseInt(str, 10, gg)
		return i
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		gg, _ := strconv.Atoi(g[4:])
		i, _ := strconv.ParseUint(str, 10, gg)
		return i
	case reflect.Float32, reflect.Float64:
		gg, _ := strconv.Atoi(g[5:])
		i, _ := strconv.ParseUint(str, 10, gg)
		return i
	case reflect.Bool:
		i,_ := strconv.ParseBool(str)
		return i
	}


	return ""
}

func Sfts() {

}

type str struct {
	A int32
	B string
}

func main() {

	var s str
	data := make(map[string][]string)
	data["A"] = []string{"1"}
	data["B"] = []string{"Bonjour"}

	Fts(data, &s)

}
