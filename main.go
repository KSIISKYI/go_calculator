package main

import (
	"fmt"

	pricePackage "project_calculator/prices"
)

func main() {
	prices := []float64{10.2, 104.3, 15.7}
	taxes := []float64{1.2, 2.3, 3.4}
	jobs := []*pricePackage.TaxIncludedPriceJob{}

	for _, taxValue := range taxes {
		job := pricePackage.NewTaxIncludedPriceJob(taxValue, prices)
		job.Process()

		jobs = append(jobs, job)
	}

	fmt.Println(jobs)
}
