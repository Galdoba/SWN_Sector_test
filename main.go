package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	temp := roll1dX(100, 0)

	fmt.Println("random =", temp)
	fmt.Println("--------------------")

	syst := NewStarSystem(16)

	fmt.Println(syst.toString())

	// race := NewVisitors()
	// visByid := VisitorByID(1)
	// fmt.Println(visByid.toString())
	// fmt.Println(race)

}
