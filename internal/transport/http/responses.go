package http

import (
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"net/http"

	"github.com/anvidev/yadom/internal/domain/barcodes"
)

func writeJSON(w http.ResponseWriter, data any, status int) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

func readJSON(w http.ResponseWriter, r *http.Request, data any) error {
	maxBytes := 1_048_578
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))
	return json.NewDecoder(r.Body).Decode(data)
}

func writeJSONError(w http.ResponseWriter, messege string, status int) error {
	type envelope struct {
		Error string `json:"error"`
	}
	return writeJSON(w, envelope{Error: messege}, status)
}

func writeImage(w http.ResponseWriter, format barcodes.ImageFormat, image image.Image) error {
	switch format {
	case barcodes.ImageFormatPNG:
		return png.Encode(w, image)
	case barcodes.ImageFormatJPEG:
		return jpeg.Encode(w, image, &jpeg.Options{Quality: 90})
	default:
		return fmt.Errorf("failed to encode image: unsupported image format")
	}
}
