package main

import (
	"fmt"
)

func main() {
	var sum , place int = -18, -18
	for i:=0; i<8; i++ {
		place += 6
		sum += place
	}
	fmt.Printf("Ans: %d\n",sum)
}