package component // import "agamigo.io/material/component"

import (
	"errors"

	"agamigo.io/gojs"
	"github.com/gopherjs/gopherjs/js"
)

// MDComponenter is the base interface for every material component
// implementation.
type MDComponenter interface {
	// MDCType should return the component's corresponding MDCType.
	MDCType() Type

	// MDCClassAttr should return the component's typical root element HTTP
	// Class Attribute. For example, an MDCCheckbox would return "mdc-checkbox".
	MDCClassAttr() string

	// SetMDC should replace a component implementation's MDComponent with mdc.
	SetMDC(mdc *C)

	// MDC should return a pointer to the component implementation's underlying
	// MDComponent. Implementors that embed a *MDComponent directly get this for
	// free.
	MDC() *C
}

type AfterStarter interface {
	// AfterStarter is implemented by components that need further setup ran
	// after their underlying MDC foundation has been initialized. rootElem is
	// provided in case you need something from the component's root
	// HTMLElement.
	AfterStart(rootElem *js.Object) error
}

// StatusType holds a component's lifecycle status.
type StatusType int

const (
	// An Uninitialized component has not been associated with the MDC library
	// yet. This package does not provide a way to access an Uninitialized
	// component.
	Uninitialized StatusType = iota

	// A Stopped component has been associated with a JS Object constructed from
	// a MDC class. New() returns a Stopped component, and Stop() will stop a
	// Running component.
	Stopped

	// A Running component has had its underlying MDC init() method called,
	// which attaches the component to a specific HTMLElement in the DOM. It is
	// ready to be used.
	Running
)

// C is the base material component type. Types that embed C and implement
// MDComponenter can use the component.Start and component.Stop functions.
type C struct {
	mdc    *js.Object
	status StatusType
}

// String returns the MDComponent's StatusType as text.
func (c *C) String() string {
	if c == nil || c.status == Uninitialized {
		return Uninitialized.String()
	}
	return c.Status().String()
}

// String returns the string version of a StatusType.
func (s StatusType) String() string {
	switch s {
	case Stopped:
		return "stopped"
	case Running:
		return "running"
	}
	return "uninitialized"
}

// Status returns the component's StatusType. For the string version use
// Status().String().
func (c *C) Status() StatusType {
	return c.status
}

func makeMDComponent(t Type) (*js.Object, error) {
	var err error
	defer gojs.CatchException(&err)
	if t == Invalid {
		return nil, errors.New("component type is Invalid")
	}

	mdcObject := js.Global.Get("mdc")

	// TODO: Move switch to component_type.go
	switch t {
	case Checkbox:
		return mdcObject.Get("checkbox").Get(t.String()), err
	case Dialog:
		return mdcObject.Get("dialog").Get(t.String()), err
	case PersistentDrawer:
		return mdcObject.Get("drawer").Get(t.String()), err
	case TemporaryDrawer:
		return mdcObject.Get("drawer").Get(t.String()), err
	case FormField:
		return mdcObject.Get("formField").Get(t.String()), err
	case GridList:
		return mdcObject.Get("gridList").Get(t.String()), err
	case IconToggle:
		return mdcObject.Get("iconToggle").Get(t.String()), err
	case LinearProgress:
		return mdcObject.Get("linearProgress").Get(t.String()), err
	case Menu:
		return mdcObject.Get("menu").Get(t.String()), err
	case Radio:
		return mdcObject.Get("radio").Get(t.String()), err
	case Ripple:
		return mdcObject.Get("ripple").Get(t.String()), err
	case Select:
		return mdcObject.Get("select").Get(t.String()), err
	case Slider:
		return mdcObject.Get("slider").Get(t.String()), err
	case Snackbar:
		return mdcObject.Get("snackbar").Get(t.String()), err
	case Tab:
		return mdcObject.Get("tabs").Get(t.String()), err
	case TabBar:
		return mdcObject.Get("tabs").Get(t.String()), err
	case TabBarScroller:
		return mdcObject.Get("tabs").Get(t.String()), err
	case TextField:
		return mdcObject.Get("textField").Get(t.String()), err
	case Toolbar:
		return mdcObject.Get("toolbar").Get(t.String()), err
	}
	return nil, err
}

// Start takes a component (c) and initializes it with an HTMLElement
// (rootElem).  The documentation for the MDComponenter and AfterStarter
// interfaces have more information.
//
// Upon success the component's status will be Running, and err will be nil.  If
// err is non-nil, it will contain any error thrown while calling the underlying
// MDC object's init() method, and the component's status will remain Stopped.
func Start(c MDComponenter, rootElem *js.Object) (err error) {
	defer gojs.CatchException(&err)

	switch {
	case rootElem == nil, rootElem == js.Undefined:
		return errors.New("rootElem is nil.")
	case c.MDC() == nil:
		c.SetMDC(&C{})
	case c.MDC().status == Running:
		return errors.New("Component already started: " +
			c.MDCType().String() + " - " + c.MDC().String())
	}

	// We create a new instance of the MDC component if MDCComponent is Stopped
	// or Uninitialized.
	newMDCClassObj, err := makeMDComponent(c.MDCType())
	if err != nil {
		return err
	}
	newMDCObj := newMDCClassObj.New(rootElem)
	c.MDC().mdc = newMDCObj
	c.MDC().status = Running

	switch co := c.(type) {
	case AfterStarter:
		err = co.AfterStart(rootElem)
		if err != nil {
			return err
		}
	}

	return err
}

// Stop stops a Running component, removing its association with an HTMLElement
// and cleaning up event listeners, etc. It changes the component's status to
// Stopped.
func Stop(mdc MDComponenter) (err error) {
	defer gojs.CatchException(&err)

	if mdc.MDC() == nil {
		return errors.New("MDC() returned nil.")
	}

	switch mdc.MDC().status {
	case Stopped:
		return errors.New("Component already stopped: " +
			mdc.MDCType().String() + " - " + mdc.MDC().String())
	case Uninitialized:
		return errors.New("Component is uninitialized: " +
			mdc.MDCType().String() + " - " + mdc.MDC().String())
	}
	mdc.MDC().mdc.Call("destroy")
	mdc.SetMDC(&C{status: Stopped})
	return err
}

// MDC implements the MDComponenter interface. Component implementations can use
// this method as-is when embedding component.C.
func (c *C) MDC() *C {
	return c
}

// MDCType implements the MDComponenter interface. This should be shadowed by a
// component implementation.
func (c *C) MDCType() Type {
	return Invalid
}

//GetObject returns the MDC component's JavaScript object.
func (c *C) GetObject() *js.Object {
	return c.mdc
}
