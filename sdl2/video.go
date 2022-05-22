package sdl2

/*
#include <SDL.h>
*/
import "C"
import "unsafe"

const (
	WINDOW_FULLSCREEN         = C.SDL_WINDOW_FULLSCREEN         // fullscreen window
	WINDOW_OPENGL             = C.SDL_WINDOW_OPENGL             // window usable with OpenGL context
	WINDOW_SHOWN              = C.SDL_WINDOW_SHOWN              // window is visible
	WINDOW_HIDDEN             = C.SDL_WINDOW_HIDDEN             // window is not visible
	WINDOW_BORDERLESS         = C.SDL_WINDOW_BORDERLESS         // no window decoration
	WINDOW_RESIZABLE          = C.SDL_WINDOW_RESIZABLE          // window can be resized
	WINDOW_MINIMIZED          = C.SDL_WINDOW_MINIMIZED          // window is minimized
	WINDOW_MAXIMIZED          = C.SDL_WINDOW_MAXIMIZED          // window is maximized
	WINDOW_INPUT_GRABBED      = C.SDL_WINDOW_INPUT_GRABBED      // window has grabbed input focus
	WINDOW_INPUT_FOCUS        = C.SDL_WINDOW_INPUT_FOCUS        // window has input focus
	WINDOW_MOUSE_FOCUS        = C.SDL_WINDOW_MOUSE_FOCUS        // window has mouse focus
	WINDOW_FULLSCREEN_DESKTOP = C.SDL_WINDOW_FULLSCREEN_DESKTOP // fullscreen window at the current desktop resolution
	WINDOW_FOREIGN            = C.SDL_WINDOW_FOREIGN            // window not created by SDL
	WINDOW_ALLOW_HIGHDPI      = C.SDL_WINDOW_ALLOW_HIGHDPI      // window should be created in high-DPI mode if supported (>= SDL 2.0.1)
	WINDOW_MOUSE_CAPTURE      = C.SDL_WINDOW_MOUSE_CAPTURE      // window has mouse captured (unrelated to INPUT_GRABBED, >= SDL 2.0.4)
	WINDOW_ALWAYS_ON_TOP      = C.SDL_WINDOW_ALWAYS_ON_TOP      // window should always be above others (X11 only, >= SDL 2.0.5)
	WINDOW_SKIP_TASKBAR       = C.SDL_WINDOW_SKIP_TASKBAR       // window should not be added to the taskbar (X11 only, >= SDL 2.0.5)
	WINDOW_UTILITY            = C.SDL_WINDOW_UTILITY            // window should be treated as a utility window (X11 only, >= SDL 2.0.5)
	WINDOW_TOOLTIP            = C.SDL_WINDOW_TOOLTIP            // window should be treated as a tooltip (X11 only, >= SDL 2.0.5)
	WINDOW_POPUP_MENU         = C.SDL_WINDOW_POPUP_MENU         // window should be treated as a popup menu (X11 only, >= SDL 2.0.5)
	WINDOW_VULKAN             = C.SDL_WINDOW_VULKAN             // window usable for Vulkan surface (>= SDL 2.0.6)
	WINDOW_METAL              = C.SDL_WINDOW_METAL              // window usable for Metal view (>= SDL 2.0.14)
)

const (
	WINDOWPOS_UNDEFINED_MASK = C.SDL_WINDOWPOS_UNDEFINED_MASK // used to indicate that you don't care what the window position is
	WINDOWPOS_UNDEFINED      = C.SDL_WINDOWPOS_UNDEFINED      // used to indicate that you don't care what the window position is
	WINDOWPOS_CENTERED_MASK  = C.SDL_WINDOWPOS_CENTERED_MASK  // used to indicate that the window position should be centered
	WINDOWPOS_CENTERED       = C.SDL_WINDOWPOS_CENTERED       // used to indicate that the window position should be centered
)

type Window C.SDL_Window

func (window *Window) cptr() *C.SDL_Window {
	return (*C.SDL_Window)(unsafe.Pointer(window))
}

func CreateWindow(title string, x, y, w, h int32, flags uint32) (*Window, error) {
	var _window = C.SDL_CreateWindow(C.CString(title), C.int(x), C.int(y), C.int(w), C.int(h), C.Uint32(flags))
	if _window == nil {
		return nil, GetError()
	}
	return (*Window)(unsafe.Pointer(_window)), nil
}

func (window *Window) Destroy() error {
	lastErr := GetError()
	ClearError()
	C.SDL_DestroyWindow(window.cptr())
	err := GetError()
	if err != nil {
		return err
	}
	SetError(lastErr)
	return nil
}

func (window *Window) GetSurface() (*Surface, error) {
	surface := (*Surface)(unsafe.Pointer(C.SDL_GetWindowSurface(window.cptr())))
	if surface == nil {
		return nil, GetError()
	}
	return surface, nil
}

func (window *Window) UpdateSurface() error {
	return errorFromInt(int(
		C.SDL_UpdateWindowSurface(window.cptr())))
}
