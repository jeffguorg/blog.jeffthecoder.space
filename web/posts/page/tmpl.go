package page

import (
	"fmt"
	"html/template"

	"github.com/jeffguorg/blog.jeffthecoder.space/internal/logging"

	tmpl "github.com/jeffguorg/blog.jeffthecoder.space/internal/template"
)

var templates *template.Template

func mustString(i interface{}, err error) string {
	if err != nil {
		panic(err)
	}
	switch i.(type) {
	case []byte:
		return string(i.([]byte))
	default:
		return fmt.Sprint(i)
	}
}

// LoadTemplates into template engine
func LoadTemplates() {
	if templates == nil {
		templates = template.New("")
		for _, filename := range tmpl.AssetNames() {
			logging.Debug("loading template: ", filename)
			if _, err := templates.New(filename).Parse(string(tmpl.MustAsset(filename))); err != nil {
				logging.Fatal(err)
			}
		}
	}
}
