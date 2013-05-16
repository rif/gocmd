package main

import (
	"errors"
	"fmt"
	"reflect"
)

type Responder struct{}

func (rs *Responder) GetCost(arg float64, reply *float64) error {
	*reply = arg * arg
	return errors.New("test")
}

func main() {
	r := new(Responder)
	v := new(float64)
	add := reflect.ValueOf(r).MethodByName("GetCost")
	ret := add.Call([]reflect.Value{reflect.ValueOf(5.0), reflect.ValueOf(v)})
	fmt.Println(*v, ret[0].Interface())
}
