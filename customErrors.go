package main

import (
	"errors"
	"fmt"
	"math"
)

type areaError struct {
	radius float64
	err string
}

type areaRectError struct {
	length float64
	width float64
	err string
}

func (e *areaError) Error() string {
	return fmt.Sprintf("radius %0.2f: %s", e.radius, e.err)
}

func (e *areaRectError) Error() string {
	return e.err
}

func (e *areaRectError) lengthNegative() bool {
	return e.length < 0
}

func (e *areaRectError) widthNegative() bool {
	return e.width < 0
}

//using new function in errors package to create a custom error
func calculateArea(radius float64) (float64, error) {
	if radius < 0 {
		return 0, errors.New("area calculation failed, negative radius")
	}
	area := radius * radius * math.Pi
	return area, nil

}

//using fmt.Errorf() to create a custom error
func calculateAreaRect(length float64, width float64) (float64, error) {
	if length < 0  {
		return 0, fmt.Errorf("Area calculation failed, %0.2f is a negative value ", length)
	}
	if width < 0  {
		return 0, fmt.Errorf("Area calculation failed, %0.2f is a negative value ", width)
	}
	return length * width, nil
}

//using struct type and it's fields to get more info about the error
func calculateArea2(radius float64) (float64, error) {
	if radius < 0 {
		return 0, &areaError{
			radius: radius,
			err:    "radius is negative",
		}
	}
	return radius * radius * math.Pi, nil
}

func calculateAreaRect2(length float64, width float64) (float64, error) {
	err := ""
	if length < 0 {
		err += "length is negative "
	}
	if width < 0 {
		err += "width is negative"
	}
	if err != "" {
		return 0, &areaRectError{
			length: length,
			width:  width,
			err:    err,
		}
	}
	return length * width, nil
}

func main() {
	//example 1 - comment out this code to see the functionality
	//radius := -19.8
	//area, err := calculateArea(radius)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println("Area ", area)

	//example 2 - comment out this code to see the functionality
	//length, width := -23.1, 9.4
	//areaRect, errorRect := calculateAreaRect(length, width)
	//if errorRect != nil {
	//	fmt.Println(errorRect)
	//	return
	//}
	//fmt.Println("Area of the rectangle is ", areaRect)

	//example 3 - comment out this code to see the functionality
	//radius2 := -3.0
	//area2, error2 := calculateArea2(radius2)
	//if error2 != nil {
	//	//using type assertion
	//	//TODO: check correction----> error2.(*areaError) means check if the concrete type of the error2 interface is areaError?
	//	if error2, ok := error2.(*areaError); ok {
	//		fmt.Printf("Radius %0.2f is negative", error2.radius)
	//		return
	//	}
	//	fmt.Println(error2)  //error than our custom areaError
	//	return
	//}
	//fmt.Println("Area is ", area2)

	//example 4
	length2, width2 := -12.0, -8.2
	areaRect2, errorRect2 := calculateAreaRect2(length2, width2)
	if errorRect2 != nil {
		if errorRect2, ok := errorRect2.(*areaRectError); ok {
			if errorRect2.lengthNegative() {
				fmt.Printf("Length %0.2f is negative", errorRect2.length)
			}
			if errorRect2.widthNegative() {
				fmt.Printf("Width %0.2f is negative", errorRect2.width)
			}
			return
		}
		fmt.Println(errorRect2)
		return
	}
	fmt.Println("area ", areaRect2)

}
