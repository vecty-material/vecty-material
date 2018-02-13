package linearprogress_test

import (
	"fmt"
	"log"

	"agamigo.io/material"
	"agamigo.io/material/linearprogress"
	"agamigo.io/material/mdctest"
	"github.com/gopherjs/gopherjs/js"
)

func Example() {
	// Create a new instance of a material linearprogress component.
	c := &linearprogress.LP{}

	// Set up a DOM HTMLElement suitable for a checkbox.
	js.Global.Get("document").Get("body").Set("innerHTML",
		mdctest.HTML(c.ComponentType().MDCClassName))
	rootElem := js.Global.Get("document").Get("body").Get("firstElementChild")

	// Start the component, which associates it with an HTMLElement.
	err := material.Start(c, rootElem)
	if err != nil {
		log.Fatalf("Unable to start component %s: %v\n", c, err.Error())
	}

	printStatus(c)
	printState(c)
	err = c.Open()
	if err != nil {
		log.Fatalf("Unable to Open component %s: %v\n", c, err.Error())
	}
	c.Determinate = false
	c.Progress = .54
	c.Buffer = 1.00
	c.Reverse = true
	err = c.Close()
	if err != nil {
		log.Fatalf("Unable to Close component %s: %v\n", c, err.Error())
	}
	printState(c)
	jsTests(c)
	printState(c)

	// Output:
	// MDCLinearProgress
	//
	// [Go] Determinate: true, Progress: 0, Buffer: 0, Reverse: false
	// [JS] Determinate: true, Progress: 0, Buffer: 0, Reverse: false
	//
	// [Go] Determinate: false, Progress: 0.54, Buffer: 1, Reverse: true
	// [JS] Determinate: false, Progress: 0.54, Buffer: 1, Reverse: true
	//
	// [Go] Determinate: true, Progress: 0.45, Buffer: 0.4, Reverse: false
	// [JS] Determinate: true, Progress: 0.45, Buffer: 0.4, Reverse: false
}

func printStatus(c *linearprogress.LP) {
	fmt.Printf("%s\n", c)
}

func printState(c *linearprogress.LP) {
	fmt.Println()
	fmt.Printf("[Go] Determinate: %v, Progress: %v, Buffer: %v, Reverse: %v\n",
		c.Determinate, c.Progress, c.Buffer, c.Reverse)
	mdcObj := c.Component().Get("foundation_")
	fmt.Printf("[JS] Determinate: %v, Progress: %v, Buffer: %v, Reverse: %v\n",
		mdcObj.Get("determinate_"),
		mdcObj.Get("progress_"),
		c.Component().Get("buffer"),
		mdcObj.Get("reverse_"),
	)
}

func jsTests(c *linearprogress.LP) {
	o := c.Component()
	o.Set("determinate", true)
	o.Set("progress", .45)
	o.Set("buffer", .40)
	o.Set("reverse", false)
}

func init() {
	// We emulate a DOM here since tests run in NodeJS.
	// Not needed when running in a browser.
	err := mdctest.Init()
	if err != nil {
		log.Fatalf("Unable to setup test environment: %v", err)
	}
}
