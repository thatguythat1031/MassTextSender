package scraper

import (
	"testing"

	"github.com/360EntSecGroup-Skylar/excelize"
)

//first step is test a fucntion to open file and print first cell: A1
//also will need to navigate to "MEN" sheet so that may need to be an option later
//looks like will likely be using excelize library

//test opening file
func TestCLI(t *testing.T) {

	t.Run("Open local .xlsx file with reader", func(t *testing.T) {
		got, err := excelize.OpenFile("Barringer.xlsx")
		want := *excelize.File{}

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
