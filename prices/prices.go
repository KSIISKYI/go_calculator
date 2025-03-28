package prices

import (
	"fmt"

	"project_calculator/helpers"
)

type TaxIncludedPriceJob struct {
	TaxRate     float64
	Prices      []float64
	TaxedPrices map[string]float64
}

func NewTaxIncludedPriceJob(taxRate float64, prices []float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate:     taxRate,
		Prices:      prices,
		TaxedPrices: make(map[string]float64),
	}
}

func (job *TaxIncludedPriceJob) String() string {
	return fmt.Sprintf("\nTaxRate: %.2f, Prices: %v, TaxedPrices: %v\n", job.TaxRate, job.Prices, job.TaxedPrices)
}

func (job *TaxIncludedPriceJob) Process() {
	for _, price := range job.Prices {
		job.TaxedPrices[fmt.Sprintf("%.2f", price)] = helpers.Round(price*job.TaxRate, 2)
	}
}
