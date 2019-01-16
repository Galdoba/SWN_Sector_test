package main

import (
	"math"
	"math/rand"
	"strconv"
	"time"
)

func FloatToString(input_num float64, roundLimit int) string {
	// to convert a float number to a string
	return strconv.FormatFloat(input_num, 'f', roundLimit, 64)
}

func randFloat(min, max float64, n int) float64 {
	res := min + rand.Float64()*(max-min)
	res = toFixed(res, 4)
	return res
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func roll1dX(x int, mod int) int {
	if x < 1 {
		x = 1
	}
	return randInt(1, x) + mod
}

func rollXdY(x int, y int) int {
	res := 0
	for i := 0; i < x; i++ {
		res = res + roll1dX(y, 0)
	}
	return res
}

func randInt(min int, max int) int {
	return min + rand.Intn(max)
}

func romanNumberStr(i int) string {
	res := ""
	switch i {
	case 1:
		res = "I"
	case 2:
		res = "II"
	case 3:
		res = "III"
	case 4:
		res = "IV"
	case 5:
		res = "V"
	case 6:
		res = "VI"
	case 7:
		res = "VII"
	case 8:
		res = "VIII"
	case 9:
		res = "IX"
	case 10:
		res = "X"
	case 11:
		res = "XI"
	case 12:
		res = "XII"
	case 13:
		res = "XIII"
	case 14:
		res = "XIV"
	case 15:
		res = "XV"
	default:
	}
	return res
}

func getRandomFromSliceStr(sl []string) string {
	l := len(sl)
	if l < 1 {
		return "Null"
	}
	return sl[randInt(0, l-1)]
}

func randomSeed() {
	rand.Seed(time.Now().UnixNano())
}

// function := map[string]func(int, int) int{
// 	"someFunction1": someFunction1,
// 	"someFunction2": someFunction2,
// }
// fmt.Println(someOtherFunction(3, 2, function["someFunction1"]))
// fmt.Println(someOtherFunction(3, 2, function["someFunction2"]))
// fmt.Println(someOtherFunction(3, 2, function["someFunction2"]))

// func someFunction1(a, b int) int {
// 	return a + b
// }

// func someFunction2(a, b int) int {
// 	return a - b
// }

// func someOtherFunction(a, b int, f func(int, int) int) int {
// 	return f(a, b)
// }
