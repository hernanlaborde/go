package main

import (
	"crypto/rand"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
)

/*
Ejemplos de ejecución.
Genera 1000 números aleatorios entre [1,7) generando archivo csv
go run cryptorandnumbers.go -cant 1000 -min 1 -max 7 -csv
Muestra el help:
go run cryptorandnumbers.go -h
*/
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

//Muy Importante = [min, max)
func main() {

	min := flag.Int64("min", 0, "Valor mínimo (limite minimo incluido [min,max) )")
	max := flag.Int64("max", 2, "Valor máximo (limite máximo excluido [min,max) )")
	cantidad := flag.Int("cant", 0, "Cantidad de números aleatorios a generar")
	csvoutput := flag.Bool("csv", false, "Genera archivo csv con los numeros aleatorios.")

	flag.Parse()

	numeros := make([]int64, 0, *cantidad)

	// csvoutput == true -> genera archivo csv
	if *csvoutput {
		filename := fmt.Sprintf("random_min-%v-max-%v_x%v.csv", *min, *max, *cantidad)
		file, err := os.Create(filename)
		if err != nil {
			message := "Error al crear el archivo csv"
			log.Fatal(message, err)
		}
		defer file.Close()
		writer := csv.NewWriter(file)
		defer writer.Flush()

		for i := 0; i < *cantidad; i++ {
			num := genRandNum(*min, *max)
			// convierte int64 en []string
			num_s := []string{strconv.FormatInt(num, 10)}
			numeros = append(numeros, num)
			if err := writer.Write(num_s); err != nil {
				log.Fatalln("Error escribiendo en el archivo csv.")
			}
		}
		fmt.Printf("Se genera archivo: %v \n\n", filename)
	} else {
		for i := 0; i < *cantidad; i++ {
			num := genRandNum(*min, *max)
			numeros = append(numeros, num)
		}
	}
	fmt.Printf("%v \n \n", numeros)
}
