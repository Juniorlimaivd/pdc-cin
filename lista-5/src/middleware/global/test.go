package global

import (
	"math/rand"
	"strconv"
)

var AccInitial = "AC"
var AccsNumber = 100

func RandomAccID() string {
	return AccInitial + strconv.Itoa(rand.Intn(AccsNumber)+1)
}

func RandomAmount() float64 {
	return float64(rand.Intn(1000)) * 100.0
}
