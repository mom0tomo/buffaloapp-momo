package actions

import (
	"html/template"
	"strings"

	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/packr"
	"github.com/gobuffalo/plush"
	"github.com/pkg/errors"
)

var r *render.Engine
var assetsBox = packr.NewBox("../public")

func init() {
	r = render.New(render.Options{
		// HTML layout to be used for all HTML requests:
		HTMLLayout: "application.html",

		// Box containing all of the templates:
		TemplatesBox: packr.NewBox("../templates"),
		AssetsBox:    assetsBox,

		// Add template helpers here:
		Helpers: render.Helpers{
			// uncomment for non-Bootstrap form helpers:
			// "form":     plush.FormHelper,
			// "form_for": plush.FormForHelper,
			"loopBeatles": func(help plush.HelperContext) (template.HTML, error) {
				if help.HasBlock() {
					s, err := help.Block()
					if err != nil {
						return "", errors.WithStack(err)
					}
					return template.HTML(strings.ToUpper(s)), nil
				}
				return "", nil
			},
		},
	})
}
