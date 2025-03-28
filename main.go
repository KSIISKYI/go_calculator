package main

import (
	"fmt"

	"project_calculator/filemanager"
	pricePackage "project_calculator/prices"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}

	prices, err := filemanager.GetPrices()

	if err != nil {
		panic("Error getting prices: " + err.Error())
	}

	taxes := []float64{1.2, 2.3, 3.4}
	jobs := []*pricePackage.TaxIncludedPriceJob{}

	for _, taxValue := range taxes {
		job := pricePackage.NewTaxIncludedPriceJob(taxValue, prices)
		job.Process()

		jobs = append(jobs, job)
	}

	fmt.Println(jobs)
}
