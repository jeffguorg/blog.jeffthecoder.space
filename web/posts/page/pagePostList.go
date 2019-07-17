package page

import (
	"net/http"

	"github.com/jeffguorg/blog.jeffthecoder.space/internal/logging"
)

func init() {
	pageMux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		if err := templates.ExecuteTemplate(w, "layout", newHelpPageData()); err != nil {
			logging.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}
