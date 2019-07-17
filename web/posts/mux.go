package posts

import (
	"github.com/go-chi/chi"
	"github.com/jeffguorg/blog.jeffthecoder.space/web/posts/api"
	"github.com/jeffguorg/blog.jeffthecoder.space/web/posts/page"
)

var (
	rootMux = chi.NewMux()
)

// Mount location into router
func Mount(router *chi.Mux, webRootPath string, paths []string) {
	api.Mount(rootMux, "/api", paths)
	page.Mount(rootMux, "/")

	router.Mount(webRootPath, rootMux)
}
