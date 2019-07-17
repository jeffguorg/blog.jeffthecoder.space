package api

import (
	"path/filepath"

	"github.com/go-chi/chi"
)

var (
	apiMux  = chi.NewMux()
	sites   map[string]string
	defSite string
)

// Mount location into router
func Mount(router *chi.Mux, location string, paths []string) {
	if len(paths) > 0 {
		sites = make(map[string]string)
		for _, path := range paths {
			dirname := filepath.Base(path)
			sites[dirname] = path
		}
		defSite = filepath.Base(paths[0])
	}
	router.Mount(location, apiMux)
}
