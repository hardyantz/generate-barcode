# barcode

create barcode generator format QR code & code 39.
using github.com/boombuler/barcode

just copy paste this file to your library

and use this library :

generateBarCode := helper.Barcode{
        Code:   "stringbarcode", 
        Width:  1000,
        Height: 160,
    }

barcode, err := generateBarCode.C39()
