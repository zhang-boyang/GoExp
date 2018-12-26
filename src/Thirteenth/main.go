package main

import (
	"fmt"
	"reflect"
)

type Rect struct {
	Width  int
	Height int
}

func main() {

	var i int = 4
	fmt.Println(reflect.ValueOf(i), reflect.TypeOf(i))
	//var inface interface{} = i
	//var in reflect.Type
	//fmt.Println(in.Kind(), in.Name(), in.NumMethod())

	type IInt int

	var in IInt = 4
	var t = reflect.TypeOf(in)
	var v = reflect.ValueOf(in)
	fmt.Println(t, v, v.Type(), v.Interface().(IInt), t.Name(), t.Size(), t.String(), t.Kind(), t.Bits())
	var R Rect = Rect{50, 50}
	var T = reflect.ValueOf(&R)
	//var _ = T.FieldByName("Width")
	//fw.SetInt(100)
	var T2 = reflect.ValueOf(&R)
	fmt.Println(reflect.TypeOf(T), reflect.TypeOf(T2.Elem()))
	var fh = T2.Elem().FieldByName("Height")
	fh.SetInt(100)
	fmt.Println(R)

}
