package main

import (
	"crypto/rand"
	"flag"
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

	min := flag.Int64("min", 0, "Valor mínimo (limite minimo incluido [min,max) )")
	max := flag.Int64("max", 2, "Valor máximo (limite máximo excluido [min,max) )")
	cantidad := flag.Int("cant", 0, "Cantidad de números aleatorios a generar")

	flag.Parse()

	for i := 0; i < *cantidad; i++ {
		num := genRandNum(*min, *max)
		fmt.Printf("%v ", num)
	}
	fmt.Println("")

}
