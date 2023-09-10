package main
import "fmt"
import "math"

type Shape interface{
	Area() float64
}

type Circle struct{
	radius float64
}

func (c Circle) Area() float64{
	return math.Pi*c.radius*c.radius
}

type Rectangle struct{
	width float64
	height float64
}

func (r Rectangle) Area() float64{
	return r.width*r.height
}

func PrintArea(s Shape){
	fmt.Printf("Area: %0.2f\n",s.Area())
}

func main(){
	var circleRadius float64
	fmt.Print("Enter the radius of the circle: ")
	fmt.Scan(&circleRadius)
	circle := Circle{radius: circleRadius}

	var rectangleWidth, rectangleHeight float64
	fmt.Print("Enter the width of the rectangle: ")
	fmt.Scan(&rectangleWidth)
	fmt.Print("Enter the height of the rectangle: ")
	fmt.Scan(&rectangleHeight)
	rectangle := Rectangle{width: rectangleWidth, height: rectangleHeight}

	PrintArea(circle)
	PrintArea(rectangle)
}

// Output
// Enter the radius of the circle: 5
// Enter the width of the rectangle: 4
// Enter the height of the rectangle: 6
// Area: 78.54
// Area: 24.00