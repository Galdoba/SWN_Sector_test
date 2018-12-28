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

	syst := NewStarSystem(temp)

	fmt.Println(syst.toString())

	race := NewVisitors()
	NewVisitors()
	NewVisitors()
	fmt.Println(race.toString())
	fmt.Println("------")

	fmt.Println(VisitorMap[0].toString())
	fmt.Println(VisitorMap[1].toString())
	fmt.Println(VisitorMap[2].toString())

}
