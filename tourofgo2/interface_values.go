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
	// 接受者也可以是nil
	if t == nil {
		fmt.Println("<nil>")
		return
	}
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
	// interface value是二元组：
	// (concrete type value, concrete type)
	var i I

	i = &T{S: "Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()

	// 接受者是 T2（注意：不是pointer）
	// 则，可以使用 T2，也可以使用 *T2 来调用
	i = T2{S: "Hello2"}
	describe(i)
	i.M()

	var ii I
	ii = &T2{S: "Hello22"}
	describe(ii)
	ii.M()

	var emptyI *T
	// concrete value本身是nil。则在nil上调用M()方法
	// 并不会跑抛空指针异常
    // interface value 是二元组 此时并不是nil哦，而是：(nil, *T)
	emptyI.M()

	// 如果是 interface value本身是nil，会咋办嗯？panic
	// interface value本是一个 二元组，如果其为nil
	// 则：interface value 即不hold concrete type value 也不hold concrete type
	var emptyI2 I2
	emptyI2.M() // panic
	fmt.Println("done")
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}


type I2 interface {
	M()
}
