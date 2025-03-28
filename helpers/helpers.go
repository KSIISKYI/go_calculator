package helpers

import (
	"math"
	"os"
	"path/filepath"
	"strconv"
)

// Round a float to a specific number of decimal places
// fix the precision issue
func Round(value float64, places int) float64 {
	shift := math.Pow(10, float64(places))

	return math.Round(value*shift) / shift
}

func GetRootPath() string {
	return os.Getenv("APP_ROOT_PATH")
}

func GetStoragePath() string {
	return filepath.Join(GetRootPath(), os.Getenv("STORAGE_DIR"))
}

func ConvertStringToFloat(value string) float64 {
	valueFloat, _ := strconv.ParseFloat(value, 64)

	return valueFloat
}
