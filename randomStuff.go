package main

import (
	"math"
	"math/rand"
	"strconv"
)

func FloatToString(input_num float64, roundLimit int) string {
	// to convert a float number to a string
	return strconv.FormatFloat(input_num, 'f', roundLimit, 64)
}

func randFloat(min, max float64, n int) float64 {
	res := min + rand.Float64()*(max-min)
	res = toFixed(res, 3)
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

func randInt(min int, max int) int {
	return min + rand.Intn(max)
}
