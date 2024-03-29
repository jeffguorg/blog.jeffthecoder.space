package web

import (
	"net/http"

	"github.com/jeffguorg/blog.jeffthecoder.space/internal/logging"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/jeffguorg/blog.jeffthecoder.space/web/mux"
	"github.com/jeffguorg/blog.jeffthecoder.space/web/posts"

	sddaemon "github.com/coreos/go-systemd/daemon"
)

// Serve starts a webserver and listen to it
func Serve(addr string, root string, paths []string, metrics bool) error {
	logging.Info("Blog root mounted to ", root)
	posts.Mount(mux.Mux(), root, paths)

	if metrics {
		logging.Info("Prometheus metrics mounted to /metrics")
		mux.Mux().Handle("/metrics", promhttp.Handler())
	}

	sddaemon.SdNotify(false, sddaemon.SdNotifyReady)
	logging.Info("Listening at ", addr)
	return http.ListenAndServe(addr, mux.Mux())
}
