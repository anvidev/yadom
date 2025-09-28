package barcodes

import "image"

type Barcode struct {
	Type   BarcodeType
	Data   string
	Width  int
	Height int
	Format ImageFormat

	// optional GS1 data - nil for regular barcodes
	GS1Data []GS1AI

	// generated image
	Image image.Image
}

type BarcodeType string

const (
	BarcodeTypeQR            BarcodeType = "qr"
	BarcodeTypeCode128       BarcodeType = "code128"
	BarcodeTypeDatamatrix    BarcodeType = "datamatrix"
	BarcodeTypeGS1Datamatrix BarcodeType = "gs1_datamatrix"
)

func (t BarcodeType) IsGS1() bool {
	switch t {
	case BarcodeTypeGS1Datamatrix:
		return true
	default:
		return false
	}
}

type ImageFormat string

const (
	ImageFormatPNG  ImageFormat = "png"
	ImageFormatJPEG ImageFormat = "jpeg"
)

func (f ImageFormat) IsValid() bool {
	switch f {
	case ImageFormatPNG, ImageFormatJPEG:
		return true
	default:
		return false
	}
}
