package main

import (
	"fmt"
)

// https://tour.golang.org/methods/6
// 入参为 pointer的 *function*，则，调用时，必须使用pointer
// pointer receiver的 *method*，则，调用时，可以pointer，可以 value
func main() {
	v := Vertex{1, 2}
	v.ScaleMethod(100) //100, 200

	p := &v
	p.ScaleMethod(99) // 9900, 19800

	fmt.Println(v)

	v1 := Vertex{1, 2}
	ScaleFunction(&v1, 9) //9, 18
	// Compile Error!
	// ScaleFunction(v1, 9) // Cannot use 'v1' (type Vertex) as type *Vertex
	fmt.Println(v1)

	fmt.Println("Begin yetAnotherScale")
	v3 := Vertex{1, 2}
	v3.yetAnotherScale(10) //11, 12
	fmt.Println(v3)        //1, 2

	v4 := Vertex{80, 90}
	p4 := &v4
	p4.yetAnotherScale(10) //90, 100
	fmt.Println(v4)       //80, 90
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
	fmt.Printf("Inside yetAnotherScale v=%v\n", v)
}
