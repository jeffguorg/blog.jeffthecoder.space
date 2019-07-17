package mux

import (
	"github.com/go-chi/chi"
	"github.com/jeffguorg/blog.jeffthecoder.space/logging"
)

var mux *chi.Mux

// Mux return a mux singleton
func Mux() *chi.Mux {
	if mux == nil {
		mux = chi.NewRouter()

		mux.Use(logging.Logger)
	}

	return mux
}
