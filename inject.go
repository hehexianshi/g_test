package main

import (
	"fmt"
	"log"
	"reflect"
)

type Inject struct {
	mapper map[reflect.Type]reflect.Value
}

func (i *Inject) Set(value interface{}) {
	i.mapper[reflect.TypeOf(value)] = reflect.ValueOf(value)
}

func (i *Inject) Get(t reflect.Type) reflect.Value {
	return i.mapper[t]
}

func (i *Inject) Invoke(inter interface{}) []reflect.Value {
	t := reflect.TypeOf(inter)
	if t.Kind() != reflect.Func {
		panic("error")
	}

	inValues := make([]reflect.Value, t.NumIn())
	for k := 0; k < t.NumIn(); k++ {
		inValues[k] = i.Get(t.In(k))
	}
	fmt.Println(inValues)

	ret := reflect.ValueOf(inter).Call(inValues)
	return ret
}

func New() *Inject {

	return &Inject{make(map[reflect.Type]reflect.Value)}
}

var inj *Inject

type LoggerInvoker func(log *log.Logger)

func (invoker LoggerInvoker) Invoker(params []interface{}) ([]reflect.Value, error) {
	invoker(params[0].(*log.Logger))
	return nil, nil
}

func Host(name string, i interface{}) {
	_, ok := i.(Inject)
	fmt.Println(ok)
	inj.Invoke(i)
}

func Logger(name string, i interface{}) {
	_, ok := i.(Inject)
	fmt.Println(ok)
	inj.Invoke(i)
}

func Dependency(a int, b string, c int) {
	fmt.Println("Dependency: ", a, b, c)
}

func main() {

	inj = New()
	inj.Set(100)
	inj.Set("abc")
	inj.Set(200)

	d := Dependency
	Host("k", d)

	Logger("t", d)
}
