package main

import (
	"os"
	"fmt"
	"time"

	"image/png"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/code39"
	"github.com/boombuler/barcode/qr"
)

// Barcode struct initialize & setup content, width, height
type Barcode struct {
	Code   string
	Width  int
	Height int
}

// StoragePath save barcode file to path
const StoragePath = "/data"

// QR generate barcode type QR
func (bc *Barcode) QR() (f string, err error) {

	qrCode, err := qr.Encode(bc.Code, qr.M, qr.Auto)

	if err != nil {
		return "", err
	}

	qrCode, err = barcode.Scale(qrCode, bc.Width, bc.Height)
	if err != nil {
		return "", err
	}

	f, err = bc.saveToFile(qrCode)

	return
}

// C39 Generate barcode type code_39
func (bc *Barcode) C39() (f string, err error) {

	c39, err := code39.Encode(bc.Code, false, true)
	b := barcode.Barcode(c39)

	c, err := barcode.Scale(b, 1000, 160)
	if err != nil {
		return "", err
	}

	f, err = bc.saveToFile(c)

	return
}

func (bc *Barcode) saveToFile(barcode barcode.Barcode) (f string, err error) {
	fileName := barcodeFileName(bc.Code)

	file, err := os.Create(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	err = png.Encode(file, barcode)
	if err != nil {
		return "", err
	}

	return fileName, nil
}

func barcodeFileName(code string) string {

	t := time.Now()
	dateNowStr := t.Format("2006/01/02/15/")

	dir := StoragePath + "barcode/" + dateNowStr

	f := code + ".png"

	os.MkdirAll(dir, 0777)
	fp := dir + f

	return fp
}

func main() {
	generateBarCode := Barcode{
		Code:   "S1447181659857",
		Width:  200,
		Height: 400,
	}

	c39, err := generateBarCode.C39()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(c39)

}
