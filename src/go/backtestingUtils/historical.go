package backtestingUtils

import (
	"github.com/luno/luno-go/decimal"
	"github.com/tealeg/xlsx"
	"os"
)

/*
TODO:
- change colNum far ask and bid
*/

//global variable for HistoricalData
var HistoricalData [][][]string

// function to process the csv file and return a 3d array of strings
// HistoricalData is of the form: [sheetNum][rowNum][colNum]
// Find more historical data files within src/go/ticker/data
// Ensure the 3rd and 4th column of every row in the file is the bid and ask price of the currency
func parseXlsx() {
	pathToHistoricalData := "../ticker/data/16to17-July.xlsx"
	fileSlice, err := xlsx.FileToSlice(pathToHistoricalData)
	if err != nil {
		panic(err)
	}
	HistoricalData = fileSlice
}

// function to get the bid price from a given row in the excel spreadsheet
func GetOfflineBid(currRow int64) decimal.Decimal {
	currPrice := HistoricalData[0][int(currRow)][3] //Change this
	// if data is non applicable skip this row
	if currPrice == "NaN" {
		return GetOfflineBid(currRow - 1)
	}

	currPriceDecimal, err := decimal.NewFromString(currPrice)

	if err != nil {
		panic(err)
	}
	return currPriceDecimal
}

// function to get the ask price from a given row in the excel spreadsheet
func GetOfflineAsk(currRow int64) decimal.Decimal {
	if currRow > 900 {
		if err := f.SaveAs("output.xlsx"); err != nil {
			panic(err)
		}
		os.Exit(3)
	}

	currPrice := HistoricalData[0][currRow][2] //Change this
	// if data is non applicable skip this row
	if currPrice == "NaN" {
		return GetOfflineAsk(currRow - 1)
	}

	currPriceDecimal, err := decimal.NewFromString(currPrice)

	if err != nil {
		panic(err)
	}

	return currPriceDecimal
}
