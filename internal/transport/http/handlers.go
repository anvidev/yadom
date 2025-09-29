package http

import (
	"net/http"

	"github.com/anvidev/yadom/internal/domain/barcodes"
)

type barcodeHandlers struct {
	service barcodes.BarcodeService
}

func NewBarcodeHandlers(service barcodes.BarcodeService) *barcodeHandlers {
	return &barcodeHandlers{service}
}

func (h barcodeHandlers) HandleListBarcodeTypes(w http.ResponseWriter, r *http.Request) {
	types := []barcodes.BarcodeType{
		barcodes.BarcodeTypeQR,
		barcodes.BarcodeTypeCode128,
		barcodes.BarcodeTypeDatamatrix,
		barcodes.BarcodeTypeGS1Datamatrix,
	}

	if err := writeJSON(w, types, http.StatusOK); err != nil {
		internalServerError(w, err)
		return
	}
}

func (h barcodeHandlers) HandleQR(w http.ResponseWriter, r *http.Request) {
	var body barcodes.BarcodeRequest

	if err := readJSON(w, r, &body); err != nil {
		badRequest(w, err)
		return
	}

	barcode, err := h.service.GenerateQR(r.Context(), body)
	if err != nil {
		internalServerError(w, err)
		return
	}

	if err := writeImage(w, barcode.Format, barcode.Image); err != nil {
		internalServerError(w, err)
		return
	}
}

func (h barcodeHandlers) HandleCode128(w http.ResponseWriter, r *http.Request) {
	var body barcodes.BarcodeRequest

	if err := readJSON(w, r, &body); err != nil {
		badRequest(w, err)
		return
	}

	barcode, err := h.service.GenerateCode128(r.Context(), body)
	if err != nil {
		internalServerError(w, err)
		return
	}

	if err := writeImage(w, barcode.Format, barcode.Image); err != nil {
		internalServerError(w, err)
		return
	}
}

func (h barcodeHandlers) HandleDatamatrix(w http.ResponseWriter, r *http.Request) {
	var body barcodes.BarcodeRequest

	if err := readJSON(w, r, &body); err != nil {
		badRequest(w, err)
		return
	}

	barcode, err := h.service.GenerateDatamatrix(r.Context(), body)
	if err != nil {
		internalServerError(w, err)
		return
	}

	if err := writeImage(w, barcode.Format, barcode.Image); err != nil {
		internalServerError(w, err)
		return
	}
}

func (h barcodeHandlers) HandleGS1Datamatrix(w http.ResponseWriter, r *http.Request) {
	var body barcodes.GS1BarcodeRequest

	if err := readJSON(w, r, &body); err != nil {
		badRequest(w, err)
		return
	}

	barcode, err := h.service.GenerateGS1Datamatrix(r.Context(), body)
	if err != nil {
		internalServerError(w, err)
		return
	}

	if err := writeImage(w, barcode.Format, barcode.Image); err != nil {
		internalServerError(w, err)
		return
	}
}
