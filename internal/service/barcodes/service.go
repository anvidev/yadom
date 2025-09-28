package barcodes

import (
	"context"
	"fmt"
	"strconv"

	"github.com/anvidev/yadom/internal/domain/barcodes"
)

type barcodeService struct {
	generator barcodes.BarcodeGenerator
}

func NewBarcodeService(generator barcodes.BarcodeGenerator) barcodes.BarcodeService {
	return &barcodeService{generator}
}

func (s barcodeService) GenerateGS1Datamatrix(ctx context.Context, req barcodes.GS1BarcodeRequest) (barcodes.Barcode, error) {
	if err := validateGS1(req.AIs); err != nil {
		return barcodes.Barcode{}, err
	}

	image, err := s.generator.GenerateGS1Datamatrix(ctx, req.AIs, req.Width, req.Height)
	if err != nil {
		return barcodes.Barcode{}, err
	}

	barcode := barcodes.Barcode{
		Type:    barcodes.BarcodeTypeGS1Datamatrix,
		Format:  req.Format,
		GS1Data: req.AIs,
		Width:   req.Width,
		Height:  req.Height,
		Image:   image,
	}

	if req.QuietZone > 0 {
		barcode.Image = s.generator.AddQuietZone(barcode.Image, req.QuietZone)
	}

	return barcode, nil
}

func (s barcodeService) GenerateQR(ctx context.Context, req barcodes.BarcodeRequest) (barcodes.Barcode, error) {
	// validate request

	image, err := s.generator.GenerateQR(ctx, req.Data, req.Width, req.Height)
	if err != nil {
		return barcodes.Barcode{}, fmt.Errorf("failed to generate qr code: %w", err)
	}

	barcode := barcodes.Barcode{
		Type:   barcodes.BarcodeTypeQR,
		Format: req.Format,
		Data:   req.Data,
		Width:  req.Width,
		Height: req.Height,
		Image:  image,
	}

	if req.QuietZone > 0 {
		barcode.Image = s.generator.AddQuietZone(barcode.Image, req.QuietZone)
	}

	return barcode, nil
}

func (s barcodeService) GenerateCode128(ctx context.Context, req barcodes.BarcodeRequest) (barcodes.Barcode, error) {
	// validate request

	image, err := s.generator.GenerateCode128(ctx, req.Data, req.Width, req.Height)
	if err != nil {
		return barcodes.Barcode{}, fmt.Errorf("failed to generate code128 barcode: %w", err)
	}

	barcode := barcodes.Barcode{
		Type:   barcodes.BarcodeTypeCode128,
		Format: req.Format,
		Data:   req.Data,
		Width:  req.Width,
		Height: req.Height,
		Image:  image,
	}

	if req.QuietZone > 0 {
		barcode.Image = s.generator.AddQuietZone(barcode.Image, req.QuietZone)
	}

	return barcode, nil
}

func (s barcodeService) GenerateDatamatrix(ctx context.Context, req barcodes.BarcodeRequest) (barcodes.Barcode, error) {
	// validate request

	image, err := s.generator.GenerateDatamatrix(ctx, req.Data, req.Width, req.Height)
	if err != nil {
		return barcodes.Barcode{}, fmt.Errorf("failed to generate code128 barcode: %w", err)
	}

	fmt.Printf("image from service is: %+v", image)

	barcode := barcodes.Barcode{
		Type:   barcodes.BarcodeTypeDatamatrix,
		Format: req.Format,
		Data:   req.Data,
		Width:  req.Width,
		Height: req.Height,
		Image:  image,
	}

	if req.QuietZone > 0 {
		barcode.Image = s.generator.AddQuietZone(barcode.Image, req.QuietZone)
	}

	return barcode, nil
}

func validateGS1(ais []barcodes.GS1AI) error {
	for i, ai := range ais {
		if !isNumeric(ai.AI) {
			return fmt.Errorf("AI at index %d must be numeric", i)
		}

		if len(ai.AI) < 2 || len(ai.AI) > 4 {
			return fmt.Errorf("AI at index %d should be 2-4 digit string", i)
		}

		if ai.Value == "" {
			return fmt.Errorf("AI %s value cannot be empty", ai.AI)
		}

		if aiInfo, found := barcodes.GetAIInfo(ai.AI); found {
			if aiInfo.Type == "fixed" && len(ai.Value) != aiInfo.Length {
				return fmt.Errorf("AI %s (%s) requires exactly %d characters, got %d",
					ai.AI, aiInfo.Name, aiInfo.Length, len(ai.Value))
			} else if aiInfo.Type == "variable" && (aiInfo.MinLength > len(ai.Value) || len(ai.Value) > aiInfo.MaxLength) {
				return fmt.Errorf("AI %s (%s) must be %d-%d characters long, got %d",
					ai.AI, aiInfo.Name, aiInfo.MinLength, aiInfo.MaxLength, len(ai.Value))
			}
		}
	}
	return nil
}

func isNumeric(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}
