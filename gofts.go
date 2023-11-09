// Change to package gofts
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

		//If the form contains a value with the same name as the struct field. Mapping must reflect value into that field.
		if o {
			//Form data mappings may contain a list of values. If the mapping has a length greater than 0.
			//Make sure that the field
			if len(g) > 0 {
				//Make sure that the field is a slice (to hold all form values)
				if d.Field(i).Kind() == reflect.Slice {

					newslice, _ := tca(g, d.Field(i).Type())
					d.Field(i).Set(reflect.ValueOf(newslice))
				}
			} else {
				//Must perform type conversion if the struct value is not of type string.
				converted, _, supported := tc(d.Field(i).Kind(), g[0])

				//Non supported fields are ignored, only supported fields are set.
				if supported {
					d.Field(i).Set(reflect.ValueOf(converted).Convert(d.Field(i).Type()))
				}

			}

		}
	}

	fmt.Print(dest)
	return nil
}

func tca(arr []string, t reflect.Type) (interface{}, error) {

	//make array slice
	//use tc to return the proper type

	return nil, nil

}

// Type conversion which returns the correct type from a given string. The form data mapping returns an array of strings.
func tc(a reflect.Kind, str string) (interface{}, error, bool) {
	//Store the string of the struct field.
	g := a.String()

	if a == reflect.String {
		return str, nil, true
	}
	switch a {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		gg, _ := strconv.Atoi(g[3:])
		i, err := strconv.ParseInt(str, 10, gg)
		return i, err, true
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		gg, _ := strconv.Atoi(g[4:])
		i, err := strconv.ParseUint(str, 10, gg)
		return i, err, true
	case reflect.Float32, reflect.Float64:
		gg, _ := strconv.Atoi(g[5:])
		i, err := strconv.ParseUint(str, 10, gg)
		return i, err, true
	case reflect.Bool:
		i, err := strconv.ParseBool(str)
		return i, err, true
	}

	return nil, nil, false
}

func Sfts() {

}

type str struct {
	A int32
	B string
	C []int
}

func main() {

	var s str
	data := make(map[string][]string)
	data["A"] = []string{"1"}
	data["B"] = []string{"Bonjour"}
	data["C"] = []string{"1", "2", "3"}

	Fts(data, &s)

}
