package scraper

import (
	"reflect"
	"testing"
)

func TestCLI(t *testing.T) {

	t.Run("extractData returns all names with their PersonInfo as a map", func(t *testing.T) {
		got, err := ExtractData("Barringer.xlsx", "Men")

		//testing 3 arbitrary entries from map and the length instead of every entry to make
		//testing less tedious
		want1 := PersonInfo{
			phone:        "(201) 742-6796",
			followUpStat: "Contacted",
			spiritSurv:   false,
			socEvent:     false,
			bStud:        false,
		}
		want2 := PersonInfo{
			phone:        "(428) 965-2160",
			followUpStat: "Attempting Contact",
			spiritSurv:   true,
			socEvent:     true,
			bStud:        true,
		}
		want3 := PersonInfo{
			phone:        "(278) 580-9114",
			followUpStat: "Uncontacted",
			spiritSurv:   false,
			socEvent:     false,
			bStud:        false,
		}

		if !reflect.DeepEqual(got["seth p"], want1) {
			t.Errorf("got %v want %v", got["seth p"], want1)
		}
		if !reflect.DeepEqual(got["test1"], want2) {
			t.Errorf("got %v want %v", got["test1"], want2)
		}
		if !reflect.DeepEqual(got["test20"], want3) {
			t.Errorf("got %v want %v", got["test20"], want3)
		}

		if err != nil {
			t.Errorf("Did not expect an error but got %s", err)
		}
	})

	t.Run("attempt to extract from non-present file", func(t *testing.T) {
		_, err := ExtractData("Miles.xlsx", "All")

		if err == nil {
			t.Errorf("Expected an error but did not get one")
		}
	})
}
