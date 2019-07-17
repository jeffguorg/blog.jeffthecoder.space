package web

import (
	"net/http"

	"github.com/jeffguorg/blog.jeffthecoder.space/web/mux"
	"github.com/jeffguorg/blog.jeffthecoder.space/web/posts"
)

// Serve starts a webserver and listen to it
func Serve(addr string, root string) error {
	posts.Mount(mux.Mux(), root)
	return http.ListenAndServe(addr, mux.Mux())
}
