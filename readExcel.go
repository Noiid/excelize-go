package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/xuri/excelize/v2"
)

func ReadExcel() {
	// f, err := excelize.OpenFile("https://api-mp.monstercode.co.id/multimedia/files/developer/document/883d4f82-0713-4083-a73e-4109288439a5.unit_block_price.xlsx")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	data, err := getData("https://api-mp.monstercode.co.id/multimedia/files/developer/document/883d4f82-0713-4083-a73e-4109288439a5.unit_block_price.xlsx")
	if err != nil {
		panic(err)
	}

	f, err := excelize.OpenReader(bytes.NewReader(data))
	if err != nil {
		fmt.Println("Reader", err)
		return
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// Get value from cell by given worksheet name and cell reference.
	cell, err := f.GetCellValue("unit_block_price", "B2")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cell)
	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("unit_block_price")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}
}

func getData(url string) ([]byte, error) {

	r, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer r.Body.Close()

	return ioutil.ReadAll(r.Body)
}
