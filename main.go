package main

import (
	"math/rand"

	"github.com/santos95mat/go-learning/fizzBuzz"
)

func main() {
	for i := 0; i < 10; i++ {
		x := rand.Intn(100)
		fizzBuzz.FizzBuzz(x)
	}
}
