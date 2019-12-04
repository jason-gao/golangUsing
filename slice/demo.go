package main

import (
	"fmt"
	"reflect"
)

func main()  {
	a := []int{1, 2, 3}
	b := []int{2, 3, 4, 5, 6}
	//a = append(a, b...)
	//fmt.Println(a)

	//a := map[string]string{"a":"1"}
	//b := map[string]string{"a":"2"}

	c := ArrayMerge(ToSlice(a), ToSlice(b))
	fmt.Println(c)
}


func ArrayMerge(ss ...[]interface{}) []interface{} {
	n := 0
	for _, v := range ss {
		n += len(v)
	}
	s := make([]interface{}, 0, n)
	for _, v := range ss {
		s = append(s, v...)
	}
	return s
}

func ToSlice(arr interface{}) []interface{} {
	v := reflect.ValueOf(arr)
	fmt.Println(v.Kind())
	if v.Kind() != reflect.Slice {
		panic("arr not slice")
	}
	l := v.Len()
	ret := make([]interface{}, l)
	for i := 0; i < l; i++ {
		ret[i] = v.Index(i).Interface()
	}
	return ret
}