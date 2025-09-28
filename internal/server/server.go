package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/anvidev/yadom/internal/infra"
	"github.com/anvidev/yadom/internal/service/barcodes"
	handlers "github.com/anvidev/yadom/internal/transport/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type api struct {
	server *http.Server
}

func New(ctx context.Context) api {
	barcodeGenerator := infra.NewBarcodeGenerator()
	barcodeService := barcodes.NewBarcodeService(barcodeGenerator)
	barcodeHandlers := handlers.NewBarcodeHandlers(barcodeService)

	r := chi.NewMux()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/v1", func(r chi.Router) {
		r.Route("/barcodes", func(r chi.Router) {
			r.Post("/qr", barcodeHandlers.HandleQR)
			r.Post("/code128", barcodeHandlers.HandleCode128)
			r.Post("/datamatrix", barcodeHandlers.HandleDatamatrix)
			r.Post("/gs1-datamatrix", barcodeHandlers.HandleGS1Datamatrix)
		})
	})

	s := &http.Server{
		Addr:    ":3900",
		Handler: r,
	}

	return api{server: s}
}

func (a api) Run() error {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := a.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("server failed to start", "error", err)
		}
	}()

	<-quit

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	return a.server.Shutdown(shutdownCtx)
}
