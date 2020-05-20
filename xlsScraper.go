package scraper

import (
	"fmt"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
)

//extractData takes a .xlsx filename (including the ".xlsx") as a string
//and returns all of the data in columns B and J.
//TODO: make sheet parameter
//TODO: make it get more info: create a struct of personinfo and store phone,
//spirit survey, soc events, bstud, and follow up status
func extractData(filename string) (map[string]string, error) {
	//open file, check for error
	file, err := excelize.OpenFile(filename)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	//create map
	people := make(map[string]string)

	i := 0

	//loop through all cells and build string with names and phone numbers
	//until empty name cell found
	for {

		//needed b/c input starts at second
		//cell in sheet.
		iConv := strconv.Itoa((i + 2))

		//get name value
		cellName, err := file.GetCellValue("Men", "B"+iConv)
		if err != nil {
			fmt.Println(err)
			return people, err
		}
		//check to make sure an entry exists, if not, stop loop
		if cellName == "" {
			break
		}

		//get phone number value
		cellNum, err := file.GetCellValue("Men", "J"+iConv)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		people[cellName] = cellNum
		i++
	}

	//if no error, return value
	return people, nil
}
