package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func readURLsFromCSV(csvFilePath string) ([][]string, error) {
	file, err := os.Open(csvFilePath)
	if err != nil {
		return nil, fmt.Errorf("the file %s does not exist", csvFilePath)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("error reading the CSV file: %v", err)
	}

	return records, nil
}

func checkPricesFromCSV(csvFilePath string) {
	records, err := readURLsFromCSV(csvFilePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	for _, row := range records {
		if len(row) != 2 {
			fmt.Printf("Error parsing the row: %v\n", row)
			continue
		}
		url := row[0]
		priceExpected, err := strconv.Atoi(row[1])
		if err != nil {
			fmt.Printf("Error parsing the price for URL %s: %v\n", url, err)
			continue
		}

		checker := NewPriceCheck(url, priceExpected)
		checker.checkPrice()
	}
}

func main() {
	csvFilePath := "price_url.csv"
	checkPricesFromCSV(csvFilePath)
}
