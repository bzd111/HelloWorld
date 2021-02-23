//go:generate statik -src=./frontend/dist -dest=./
package main

import (
	"fmt"
	"net/http"
	"os"

	_ "static-embed/statik"

	"github.com/rakyll/statik/fs"
	"go.uber.org/zap"
)

func main() {
	statikFS, err := fs.New()
	if err != nil {
		zap.L().Error("", zap.Error(err))
		os.Exit(1)
	}
	http.Handle("/", http.StripPrefix("/", http.FileServer(statikFS)))
	err = http.ListenAndServe(fmt.Sprintf(":%d", 8080), nil)
	if err != nil {
		zap.L().Error("", zap.Error(err))
		os.Exit(1)
	}
}
