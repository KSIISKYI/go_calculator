package main

import "fmt"

func main() {
	prices := []float64{10.2, 104.3, 15.7}
	taxes := []float64{1.2, 2.3, 3.4}

	taxedPrices := map[float64][]float64{}

	for _, taxValue := range taxes {
		slice := []float64{}

		for _, price := range prices {
			slice = append(slice, price*taxValue)
		}

		taxedPrices[taxValue] = slice
	}

	fmt.Println(taxedPrices)
}
