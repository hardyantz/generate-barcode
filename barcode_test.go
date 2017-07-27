package main

import (
	"testing"
)

func TestBarcode_QR(t *testing.T) {
	bc := Barcode{
		Code:"abcde",
		Width: 1000,
		Height: 160,
	}

	name := barcodeFileName(bc.Code)
	expectedResult := name

	results, _ := bc.QR()

	if results != expectedResult {
		t.Errorf("Test failed, expected: %s, \n results: %s",
		expectedResult, results)
	}
}

func TestBarcode_C39(t *testing.T) {
	bc := Barcode{
		Code:"s1223332",
		Width: 1000,
		Height: 160,
	}

	name := barcodeFileName(bc.Code)
	expectedResult := name

	results, _ := bc.C39()

	if results != expectedResult {
		t.Errorf("Test failed, expected: %s, \n results: %s",
			expectedResult, results)
	}
}


