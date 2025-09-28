package barcodes

import "fmt"

var (
	ErrInvalidBarcodeType error = fmt.Errorf("invalid barcode type")
)
