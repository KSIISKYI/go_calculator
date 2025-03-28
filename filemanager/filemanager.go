package filemanager

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"project_calculator/helpers"
)

const (
	pricesFile = "prices.json"
)

func GetPrices() ([]float64, error) {
	jsonFile, err := getPricesFile()

	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}

	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)
	prices, err := decodePricesRaw(byteValue)

	if err != nil {
		fmt.Println("Error decoding prices:", err)
	}

	return prices, err
}

func decodePricesRaw(raw []byte) (prices []float64, err error) {
	rawPrices := []interface{}{}

	if len(raw) != 0 {
		err = json.Unmarshal(raw, &rawPrices)
	}

	for _, price := range rawPrices {
		priceNumber, isNumber := price.(float64)

		if !isNumber {
			priceNumber = helpers.ConvertStringToFloat(price.(string))
		}

		prices = append(prices, priceNumber)
	}

	return prices, nil
}

func getPricesFile() (*os.File, error) {
	filePath := filepath.Join(helpers.GetStoragePath(), pricesFile)

	return os.Open(filePath)
}
