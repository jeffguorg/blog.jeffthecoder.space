package posts

import (
	"github.com/go-chi/chi"
)

var rootMux = chi.NewMux()

// Mount location into router
func Mount(router *chi.Mux, webRootPath string) {
	router.Handle(webRootPath, rootMux)
}
