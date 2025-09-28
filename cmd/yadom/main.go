package main

import (
	"context"

	"github.com/anvidev/yadom/internal/server"
)

func main() {
	s := server.New(context.Background())
	s.Run()
}
