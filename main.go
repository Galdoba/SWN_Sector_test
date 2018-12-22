package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	temp := roll1dX(100, 0)
	for temp > 3 {
		temp = roll1dX(100, 0)
	}
	fmt.Println("random =", temp)
	fmt.Println("--------------------")
	for i := 25; i > 20; i-- {
		syst := NewStarSystem(i)
		fmt.Println("TEST CASE", i)
		fmt.Println(syst.toString())
		fmt.Println("\n\n\n")
	}

}
