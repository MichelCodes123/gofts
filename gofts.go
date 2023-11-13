// Change to package gofts
package gofts

import (
	"fmt"
	"reflect"
	"strconv"
)

// Transfers form data into the struct specified by the user. Returns false if the operation fails, and error message for invalid inputs
// Throws error if:
// The dest is not a pointer to a struct, type conversion cannot be made
func Fts(form map[string][]string, dest interface{}) error {

	v := reflect.ValueOf(dest)
	if v.Kind() != reflect.Ptr {
		return fmt.Errorf("%s", "Must be a pointer to a struct")
	}

	d := v.Elem()
	if d.Kind() != reflect.Struct {
		return fmt.Errorf("%s", "Not value of struct")
	}

	for i := 0; i < d.NumField(); i++ {
		
		nameOfStructField := d.Type().Field(i).Name
		g, o := form[nameOfStructField]

		if o {
			//Form data mappings may contain a list of values. Check if this is the case
			if len(g) > 1 {
				//Ensure that the field can support the list of values... Should be kind of slice.
				if d.Field(i).Kind() == reflect.Slice {
					newslice, err, supported := type_convert_slice(g, d.Field(i).Type())
					if err != nil {
						return err
					}
					if supported {
						d.Field(i).Set(newslice)
					}
				}
			} else {
		
				converted, err, supported := type_convert(d.Field(i).Kind(), g[0])
				if err != nil {
					return err
				}
				//Non supported fields are ignored, only supported fields are set.
				if supported {
					d.Field(i).Set(reflect.ValueOf(converted).Convert(d.Field(i).Type()))
				}

			}
		}
	}

	return nil
}

func type_convert_slice(arr []string, t reflect.Type) (reflect.Value, error, bool) {
	//t is the array type. I.e []int, []float32
	//t.Elem() returns the type of an element. I.e int
	//t.Elem.kind() returns the underlying type of the element. I.e int
	//Make sure of this ^

	switch t.Elem().Kind() {
	case reflect.String, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64, reflect.Bool:
		n := reflect.MakeSlice(t, 0, len(arr))
		for _, v := range arr {
			converted, err, _ := type_convert(t.Elem().Kind(), v)
			if err != nil {
				return reflect.ValueOf("0"), err, false
			}
			n = reflect.Append(n, reflect.ValueOf(converted).Convert(t.Elem()))
		}

		return n, nil,true
	}

	return reflect.ValueOf("0"), nil, false

}

// Type conversion which returns the correct type from a given string. The form data mapping returns an array of strings.
func type_convert(a reflect.Kind, str string) (interface{}, error, bool) {
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

func Mfts() {

}
