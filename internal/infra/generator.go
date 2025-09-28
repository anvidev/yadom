package infra

import (
	"bytes"
	"context"
	"image"
	"image/color"

	"github.com/anvidev/yadom/internal/domain/barcodes"
	bc "github.com/boombuler/barcode"
	"github.com/boombuler/barcode/code128"
	"github.com/boombuler/barcode/datamatrix"
	"github.com/boombuler/barcode/qr"
)

type boombulerGenerator struct{}

func NewBarcodeGenerator() barcodes.BarcodeGenerator {
	return &boombulerGenerator{}
}

func (g boombulerGenerator) GenerateGS1Datamatrix(ctx context.Context, ais []barcodes.GS1AI, width, height int) (image.Image, error) {
	gs1String := BuildGS1DataString(ais)

	barcode, err := datamatrix.Encode(gs1String)
	if err != nil {
		return nil, err
	}

	barcode, err = bc.Scale(barcode, width, height)
	if err != nil {
		return nil, err
	}

	return barcode, nil
}

func (g boombulerGenerator) GenerateQR(ctx context.Context, data string, width, height int) (image.Image, error) {
	barcode, err := qr.Encode(data, qr.M, qr.Auto)
	if err != nil {
		return nil, err
	}

	barcode, err = bc.Scale(barcode, width, height)
	if err != nil {
		return nil, err
	}

	return barcode, nil
}

func (g boombulerGenerator) GenerateCode128(ctx context.Context, data string, width, height int) (image.Image, error) {
	image, err := code128.Encode(data)
	if err != nil {
		return nil, err
	}

	scaled, err := bc.Scale(image, image.Bounds().Dx(), height)
	if err != nil {
		return nil, err
	}

	return scaled, nil
}

func (g boombulerGenerator) GenerateDatamatrix(ctx context.Context, data string, width, height int) (image.Image, error) {
	image, err := datamatrix.Encode(data)
	if err != nil {
		return nil, err
	}

	image, err = bc.Scale(image, width, height)
	if err != nil {
		return nil, err
	}

	return image, nil
}

func (g *boombulerGenerator) AddQuietZone(img image.Image, width int) image.Image {
	return quietZone{width: width, barcode: img}
}

type quietZone struct {
	width   int
	barcode image.Image
}

func (z quietZone) At(x, y int) color.Color {
	bounds := z.barcode.Bounds()
	w := bounds.Dx()
	h := bounds.Dy()
	if x < z.width || x >= w+z.width || y < z.width || y >= h+z.width {
		return color.White
	}
	return z.barcode.At(x-z.width, y-z.width)
}

func (z quietZone) Bounds() image.Rectangle {
	b := z.barcode.Bounds()
	return image.Rect(0, 0, b.Dx()+2*z.width, b.Dy()+2*z.width)
}

func (z quietZone) ColorModel() color.Model {
	return z.barcode.ColorModel()
}

func BuildGS1DataString(ais []barcodes.GS1AI) string {
	parts := new(bytes.Buffer)
	parts.WriteByte(barcodes.FNC1)

	for i, ai := range ais {
		part := new(bytes.Buffer)
		part.WriteString(ai.AI)
		part.WriteString(ai.Value)

		if i < len(ais)-1 {
			if barcodes.IsVariableLengthAI(ai.AI) {
				part.WriteByte(barcodes.FNC1)
			}
		}

		parts.Write(part.Bytes())
	}

	return parts.String()
}
