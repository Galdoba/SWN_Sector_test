package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	temp := roll1dX(100, 0)
	temp = 3
	fmt.Println("random =", temp)
	fmt.Println("--------------------")

	syst := NewStarSystem(temp)
	fmt.Println(syst.toString())
}
