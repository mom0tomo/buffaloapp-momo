package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/pkg/errors"
)

type Band struct {
	Name string
	Bio  string
}

func BandsNew(c buffalo.Context) error {
	c.Set("band", Band{})
	return c.Render(200, r.HTML("bands/new.html"))
}

func BandsCreate(c buffalo.Context) error {
	b := &Band{}
	if err := c.Bind(b); err != nil {
		return errors.WithStack(err)
	}
	c.Set("band", b)
	return c.Render(201, r.HTML("bands/show.html"))
}
