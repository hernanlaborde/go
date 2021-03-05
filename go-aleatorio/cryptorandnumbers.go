package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func genRandNum(min, max int64) int64 {
	//Como rand.Int solo entrega números de entre [0,max) se resta a max
	dif := big.NewInt(max - min)

	//Int returns a uniform random value in [0, max). It panics if max <= 0.
	n, err := rand.Int(rand.Reader, dif)
	if err != nil {
		panic(err)
	}

	//se vuelve a sumar el minimo para neutralizar la resta del inicio.
	return n.Int64() + min
}

// Muy Importante = [min, max)
func main() {
	var min, max int64
	var cantidad int

	min = 0
	max = 101
	cantidad = 10

	for i := 0; i < cantidad; i++ {
		num := genRandNum(min, max)
		fmt.Printf("%v ", num)
	}
	fmt.Println("")

}
