package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/mom0tomo/myapp/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
