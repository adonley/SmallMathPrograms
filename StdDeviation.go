package main

import (
	"fmt"
)

func main() {
	var one,two float64 = calcStdDeviation()
	fmt.Printf("Avg: %f, Variance: %f\n",one,two)
}

func calcStdDeviation() (ans, avg float64) {
	var input int
	var aggrigate, aggrigate2, input2 float64 = 0,0,0
	fmt.Printf("%s\n","Calculate Std Deviation.")
	fmt.Printf("%s\n","------------------------")
	fmt.Printf("%s\n","How many numbers?")
	fmt.Scanf("%d",&input)
	for i:=0; i < input; i++ {
		fmt.Printf("%d: ",i)
		fmt.Scanf("%f",&input2)
		aggrigate += input2
		aggrigate2 += input2*input2
	}
	avg = aggrigate/float64(input)
	ans = aggrigate/float64(input)
	return ans, avg
}