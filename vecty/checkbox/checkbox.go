package checkbox

import (
	"time"

	mdccheckbox "agamigo.io/material-components-go/mdc/checkbox"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

type C struct {
	*mdccheckbox.C
	vecty.Core
}

func New() *C {
	c := &C{}
	c.C = mdccheckbox.New()
	return c
}

func (c *C) Render() vecty.ComponentOrHTML {
	println("Render checkbox called")
	e := elem.Div(
		vecty.Markup(
			vecty.Class("mdc-checkbox"),
		),
		elem.Input(
			vecty.Markup(
				prop.Type(prop.TypeCheckbox),
				vecty.Class("mdc-checkbox__native-control"),
				prop.ID("native-js-checkbox"),
			),
		),
		elem.Div(
			vecty.Markup(
				vecty.Class("mdc-checkbox__background"),
				vecty.UnsafeHTML(
					`<svg class="mdc-checkbox__checkmark"
						viewBox="0 0 24 24">
					<path class="mdc-checkbox__checkmark__path"
						fill="none"
						stroke="white"
						d="M1.73,12.91 8.1,19.28 22.79,4.59"/>
					</svg>`,
				),
			),
			elem.Div(
				vecty.Markup(
					vecty.Class("mdc-checkbox__mixedmark"),
				),
			),
		),
	)
	return e
}

func (c *C) Mount() {
	println("Mount checkbox called")
	c.C.Start()
	go c.testCB()
}

func (c *C) Unmount() {
	println("Unmount checkbox called")
	c.C.Stop()
}

func (c *C) testCB() {
	for _ = range time.Tick(1 * time.Second) {
		s := c.State()
		print(s)
		if s == mdccheckbox.INDETERMINATE_DISABLED {
			c.SetState(mdccheckbox.UNCHECKED)
			continue
		}
		c.SetState(s + mdccheckbox.DISABLED)
	}
}
