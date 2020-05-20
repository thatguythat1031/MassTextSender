package scraper

import (
	"reflect"
	"testing"
)

//first step is test a fucntion to open file and print first cell: A1
//also will need to navigate to "MEN" sheet so that may need to be an option later
//looks like will likely be using excelize library

//test opening file
func TestCLI(t *testing.T) {

	t.Run("extractData returns first all names with their phone numbers", func(t *testing.T) {
		got, err := extractData("Barringer.xlsx")
		want := map[string]string{
			"seth p":  "(201) 742-6796",
			"chris c": "(984) 938-7156",
			"fiona r": "(354) 388-5157",
			"paul o":  "(893) 828-6784",
			"meg s":   "(281) 658-2936",
			"baily v": "(841) 508-0541",
			"test1":   "(428) 965-2160",
			"test2":   "(575) 770-5630",
			"test3":   "(947) 472-4409",
			"test4":   "(617) 670-6033",
			"test5":   "(724) 992-5937",
			"test6":   "(385) 494-5408",
			"test7":   "(646) 677-9739",
			"test8":   "(263) 262-5372",
			"test9":   "(610) 213-4025",
			"test10":  "(946) 623-8615",
			"test11":  "(499) 881-3999",
			"test12":  "(963) 315-7080",
			"test13":  "(996) 656-4599",
			"test14":  "(721) 271-5887",
			"test15":  "(516) 458-8837",
			"test16":  "(600) 687-5571",
			"test17":  "(688) 992-4106",
			"test18":  "(290) 745-1183",
			"test19":  "(510) 224-7039",
			"test20":  "(278) 580-9114",
			"test21":  "(509) 991-5196",
			"test22":  "(810) 572-7101",
			"test23":  "(783) 773-6362",
			"test24":  "(687) 350-8556",
			"test25":  "(504) 542-5459",
			"test26":  "(643) 817-2536"}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %s want %s", got, want)
		}

		if err != nil {
			t.Errorf("Did not expect an error but got %s", err)
		}
	})
}
