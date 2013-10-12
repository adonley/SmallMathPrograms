package main

import (
	"fmt"
    "math"
)

func herons_function(a, b, c float64) (float64) {
	s:= (a+b+c)/2
	temp:=s*(s-a)*(s-b)*(s-c)
	s = math.Sqrt(temp)
	return  s
}

func main() {
	var a,b,c float64
	fmt.Printf("a: ")
	fmt.Scanf("%f",&a)
	fmt.Printf("b: ")
	fmt.Scanf("%f",&b)
	fmt.Printf("c: ")
	fmt.Scanf("%f",&c)
	fmt.Printf("%s Ans:%f\n","Heron's Formula ", herons_function(a,b,c))
}