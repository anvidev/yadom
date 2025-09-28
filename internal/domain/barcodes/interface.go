package barcodes

import (
	"context"
	"image"
)

type BarcodeRequest struct {
	Format ImageFormat `json:"imageFormat"`

	Width     int `json:"width"`
	Height    int `json:"height"`
	QuietZone int `json:"quietZone"`

	Data string `json:"data"`
}

type GS1BarcodeRequest struct {
	Format ImageFormat `json:"imageFormat"`

	Width     int `json:"width"`
	Height    int `json:"height"`
	QuietZone int `json:"quietZone"`

	AIs []GS1AI `json:"ais"`
}

type BarcodeService interface {
	GenerateQR(ctx context.Context, req BarcodeRequest) (Barcode, error)
	GenerateCode128(ctx context.Context, req BarcodeRequest) (Barcode, error)
	GenerateDatamatrix(ctx context.Context, req BarcodeRequest) (Barcode, error)
	GenerateGS1Datamatrix(ctx context.Context, req GS1BarcodeRequest) (Barcode, error)
}

type BarcodeGenerator interface {
	GenerateQR(ctx context.Context, data string, width, height int) (image.Image, error)
	GenerateCode128(ctx context.Context, data string, width, height int) (image.Image, error)
	GenerateDatamatrix(ctx context.Context, data string, width, height int) (image.Image, error)
	GenerateGS1Datamatrix(ctx context.Context, ais []GS1AI, width, height int) (image.Image, error)
	AddQuietZone(image image.Image, width int) image.Image
}
