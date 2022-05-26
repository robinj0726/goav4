package sdl2

/*
#include <SDL.h>
*/
import "C"
import "unsafe"

type Surface struct {
	flags    uint32         // (internal use)
	Format   *PixelFormat   // the format of the pixels stored in the surface (read-only) (https://wiki.libsdl.org/SDL_PixelFormat)
	W        int32          // the width in pixels (read-only)
	H        int32          // the height in pixels (read-only)
	Pitch    int32          // the length of a row of pixels in bytes (read-only)
	pixels   unsafe.Pointer // the pointer to the actual pixel data; use Pixels() for access
	UserData unsafe.Pointer // an arbitrary pointer you can set
	locked   int32          // used for surfaces that require locking (internal use)
	lockData unsafe.Pointer // used for surfaces that require locking (internal use)
	ClipRect Rect           // a Rect structure used to clip blits to the surface which can be set by SetClipRect() (read-only)
	_        unsafe.Pointer // map; info for fast blit mapping to other surfaces (internal use)
	RefCount int32          // reference count that can be incremented by the application
}

func (surface *Surface) cptr() *C.SDL_Surface {
	return (*C.SDL_Surface)(unsafe.Pointer(surface))
}

func (surface *Surface) FillRect(rect *Rect, color uint32) error {
	if C.SDL_FillRect(surface.cptr(), (*C.SDL_Rect)(unsafe.Pointer(rect)), C.Uint32(color)) != 0 {
		return GetError()
	}
	return nil
}
