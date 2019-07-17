package mux

import (
	"github.com/go-chi/chi"
	"github.com/jeffguorg/blog.jeffthecoder.space/logging"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var mux *chi.Mux

// Mux return a mux singleton
func Mux() *chi.Mux {
	if mux == nil {
		mux = chi.NewRouter()

		mux.Use(logging.Logger)

		mux.Handle("/metrics", promhttp.Handler())
	}

	return mux
}
