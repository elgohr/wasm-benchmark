package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	value := 0
	numberOfPrimes := 0
	sequencesSinceStart := 0
	second := time.Now().Second()
	for sequencesSinceStart <= 10 {
		currentSecond := time.Now().Second()
		if second != currentSecond {
			println(fmt.Sprintf("%v-%v", sequencesSinceStart, numberOfPrimes))
			second = currentSecond
			numberOfPrimes = 0
			sequencesSinceStart++
		}
		if IsPrime(value) {
			numberOfPrimes++
		}
		value++
	}
}

func IsPrime(value int) bool {
	for i := 2; i <= int(math.Floor(float64(value)/2)); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}
