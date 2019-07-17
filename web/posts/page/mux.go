package page

import (
	"github.com/go-chi/chi"
)

var (
	pageMux = chi.NewMux()
	sites   []string
	defSite string
)

// Mount location into router
func Mount(router *chi.Mux, path string) {
	if templates == nil {
		LoadTemplates()
	}

	router.Mount(path, pageMux)
}
