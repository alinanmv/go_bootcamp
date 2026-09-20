package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const file = "test.txt"

func GetFloatFromFile() (float64, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return 1000, errors.New("no file was found")
	}
	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)
	if err != nil {
		return 1000, errors.New("can't convert to a float")
	}
	return value, nil
}
func WriteFloatToFile(value float64) {
	valueText := fmt.Sprint(value)
	os.WriteFile(file, []byte(valueText), 0644)
}
