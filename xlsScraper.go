package scraper

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
)

//PersonInfo holds all of the information that needs to be extracted from the
//.xlsx file. Follow up status should be "Contacted", "Uncontacted", or "Attempting Contact"
type PersonInfo struct {
	phone        string
	followUpStat string
	spiritSurv   bool
	socEvent     bool
	bStud        bool
}

//ExtractData takes a .xlsx filename (including the ".xlsx") as a string
//and returns all of the data in columns B and J.
//TODO: make it get more info: create a struct of personinfo and store phone,
//spirit survey, soc events, bstud, and follow up status
func ExtractData(filename string, sheetname string) (map[string]PersonInfo, error) {
	//open file, check for error
	file, err := excelize.OpenFile(filename)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	//create map
	people := make(map[string]PersonInfo)

	i := 0

	//loop through all cells and build string with names and phone numbers
	//until empty name cell found
	for {

		//needed b/c input starts at second
		//cell in sheet.
		iConv := strconv.Itoa((i + 2))

		//get name value
		cellName, err := file.GetCellValue(sheetname, "B"+iConv)
		if err != nil {
			fmt.Println(err)
			return people, err
		}
		//check to make sure an entry exists, if not, stop loop
		if cellName == "" {
			break
		}

		//get phone number value
		cellNumPhone, err := file.GetCellValue(sheetname, "J"+iConv)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		//get followUpStat value
		cellFollowUp, err := file.GetCellValue(sheetname, "P"+iConv)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		//get spiritSurv value
		cellSpiritSurv, err := file.GetCellValue(sheetname, "L"+iConv)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		spiritSurvVal := convertCellStringToBool(cellSpiritSurv)

		//get socEvent value
		cellSocEvent, err := file.GetCellValue(sheetname, "L"+iConv)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		socEventVal := convertCellStringToBool(cellSocEvent)

		//get bstud value
		cellBStud, err := file.GetCellValue(sheetname, "L"+iConv)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		bStudVal := convertCellStringToBool(cellBStud)

		people[cellName] = PersonInfo{
			phone:        cellNumPhone,
			followUpStat: cellFollowUp,
			spiritSurv:   spiritSurvVal,
			socEvent:     socEventVal,
			bStud:        bStudVal,
		}
		i++
	}

	//if no error, return map
	return people, nil
}

//Takes the given cell value and converts it to a bool for use in
//the PersonInfo struct
func convertCellStringToBool(cellString string) bool {
	var convertedBool bool

	if cellString == "" || strings.EqualFold("no", cellString) {
		convertedBool = false
	} else { //if cell is "maybe" or "yes"
		convertedBool = true
	}

	return convertedBool
}
