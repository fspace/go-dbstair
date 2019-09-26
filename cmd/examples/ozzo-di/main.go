package main

import (
	"fmt"
	"github.com/go-ozzo/ozzo-di"
	"reflect"
)

type Bar interface {
	String() string
}

func test(bar Bar) {
	fmt.Println(bar.String())
}

type Foo struct {
	s string
}

func (f *Foo) String() string {
	return f.s
}

type MyBar struct {
	Bar `inject`
}

func main() {
	// creating a DI container
	c := di.NewContainer()

	// register a Foo instance as the Bar interface type
	c.RegisterAs(&Foo{"hello"}, di.InterfaceOf((*Bar)(nil)))

	// &Foo{"hello"} will be injected as the Bar parameter for test()
	c.Call(test)
	// Output:
	// hello

	// create a MyBar object and inject its Bar field
	bar := c.Make(reflect.TypeOf(&MyBar{})).(Bar)
	fmt.Println(bar.String())
	// Output:
	// hello
}
