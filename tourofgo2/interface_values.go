package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

type T2 struct {
	S string
}

type F float64

func (t *T) M() {
	fmt.Println(t.S)
}

func (t T2) M() {
	fmt.Println(t.S)
}

func (f F) M() {
	fmt.Println(f)
}

func main() {
	// i就是interface value
	// 可以认为 interface value 是二元组：
	// (下面的 特定concrete type的value, interface concrete type)
	var i I

	i = &T{S: "Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()

	i = T2{S: "Hello2"}
	describe(i)
	i.M()

	i = &T2{S: "Hello22"}
	describe(i)
	i.M()

}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
