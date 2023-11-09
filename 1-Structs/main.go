package main

import (
	"fmt"
)

type Inner struct {
	InnerField int
}

func (i *Inner) InnerMethod() {
	fmt.Println("InnerMethod called with InnerField value:", i.InnerField)
}

func (i *Inner) Overwrite() {
	fmt.Println("Inner Overwrite")
}

type Outer struct {
	Inner
	OuterField int
}

func (o *Outer) OuterMethod() {
	fmt.Println("OuterMethod called with OuterField value:", o.OuterField)
}

func (o *Outer) Overwrite() {
	fmt.Println("Outer Overwrite")
}

func main() {
	outerInstance := Outer{
		Inner: Inner{
			InnerField: 42,
		},
		OuterField: 24,
	}

	outerInstance.InnerMethod()
	outerInstance.OuterMethod()
	outerInstance.Overwrite()
}
