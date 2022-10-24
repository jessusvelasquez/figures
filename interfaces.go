package figures

import "fmt"

type geometry interface {
	getArea() float32
	getPerimeter() float32
}

func F_figures(ifg geometry) {
	fmt.Println("Figure: ", ifg)
	fmt.Println("Area: ", ifg.getArea())
	fmt.Println("Perimeter: ", ifg.getPerimeter())
}
