package main

import (
	"fmt"
)

// https://tour.golang.org/methods/6
// 入参为 pointer的 *function*，则，调用时，必须使用pointer
// pointer receiver的 *method*，则，调用时，可以pointer，可以 value
func main() {
	v := Vertex{5, 6}
	v.ScaleMethod(100)

	p := &v
	p.ScaleMethod(99)

	fmt.Println(v)

	v1 := Vertex{1, 2}
	ScaleFunction(&v1, 9)
	//ScaleFunction(v1, 9) // Compile Error!
	fmt.Println(v1)

	v3 := Vertex{80, 90}
	v3.yetAnotherScale(12) //92, 102
	fmt.Println(v3)        //80, 90

	v4 := Vertex{11, 22}
	p4 := &v4
	p4.yetAnotherScale(2) //13, 24
	fmt.Println(v4)       //11, 22
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) ScaleMethod(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunction(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v Vertex) yetAnotherScale(f float64) {
	v.X = v.X + f
	v.Y = v.Y + f
	fmt.Println(v)
}
