package main

import (
	"fmt"
	"reflect"
)

func Hello() string {
	return "hi i've just imported some packages"

}

// Transfers form data into the struct specified by the user. Returns false if the operation fails, and error message for invalid inputs
//Throws error if: 
//The dest is not a pointer to a struct
func Fts(form map[string][]string, dest interface{})  error {

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
			if (len(g) > 0){

			}
			d.Field(i).SetString(g[0])
		}

	}

	fmt.Print(dest)
	return nil
}

func Sfts() {

}

type str struct {
	A string
	B string
}

func main() {

	var s str
	data := make(map[string][]string)
	data["A"] = []string{"hello"}
	data["B"] = []string{"Bonjour"}

	fmt.Print(data["g"])

	Fts(data, &s)

}
