package imgui

import (
	"errors"
)

// #include "wrapper/Context.h"
// #include "plot_wrapper/Contexts.h"
import "C"

// Context specifies a scope of ImGui/ImPlot.
//
// All contexts share a same FontAtlas by default.
// If you want different font atlas, you can create them and overwrite the CurrentIO.Fonts of an ImGui context.
type Context struct {
	handle C.IggContext
	plot   C.igpContext
}

// CreateContext produces a new internal state scope.
// Passing nil for the fontAtlas creates a default font.
func CreateContext(fontAtlas *FontAtlas) *Context {
	var fontAtlasPtr C.IggFontAtlas
	if fontAtlas != nil {
		fontAtlasPtr = fontAtlas.handle()
	}
	return &Context{handle: C.iggCreateContext(fontAtlasPtr), plot: C.igpCreateContext()}
}

// ErrNoContext is used when no context is current.
var ErrNoContext = errors.New("no current context")

// CurrentContext returns the currently active state scope.
// Returns ErrNoContext if no context is available.
func CurrentContext() (*Context, error) {
	raw := C.iggGetCurrentContext()
	if raw == nil {
		return nil, ErrNoContext
	}
	return &Context{handle: raw, plot: C.igpCurrentContext()}, nil
}

// Destroy removes the internal state scope.
// Trying to destroy an already destroyed context does nothing.
func (context *Context) Destroy() {
	if context.handle != nil {
		C.iggDestroyContext(context.handle)
		context.handle = nil
	}
	if context.plot != nil {
		C.igpDestroyContext(context.plot)
		context.plot = nil
	}
}

// ErrContextDestroyed is returned when trying to use an already destroyed context.
var ErrContextDestroyed = errors.New("context is destroyed")

// SetCurrent activates this context as the currently active state scope.
func (context Context) SetCurrent() error {
	if context.handle == nil {
		return ErrContextDestroyed
	}
	C.iggSetCurrentContext(context.handle)
	C.igpSetCurrentContext(context.plot)
	return nil
}
