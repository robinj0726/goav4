package sdl2

/*
#include <SDL.h>
*/
import "C"
import "unsafe"

const (
	RENDERER_SOFTWARE      = C.SDL_RENDERER_SOFTWARE      // the renderer is a software fallback
	RENDERER_ACCELERATED   = C.SDL_RENDERER_ACCELERATED   // the renderer uses hardware acceleration
	RENDERER_PRESENTVSYNC  = C.SDL_RENDERER_PRESENTVSYNC  // present is synchronized with the refresh rate
	RENDERER_TARGETTEXTURE = C.SDL_RENDERER_TARGETTEXTURE // the renderer supports rendering to texture
)

const (
	TEXTUREACCESS_STATIC    = C.SDL_TEXTUREACCESS_STATIC    // changes rarely, not lockable
	TEXTUREACCESS_STREAMING = C.SDL_TEXTUREACCESS_STREAMING // changes frequently, lockable
	TEXTUREACCESS_TARGET    = C.SDL_TEXTUREACCESS_TARGET    // can be used as a render target
)

type Renderer C.SDL_Renderer

func (r *Renderer) cptr() *C.SDL_Renderer {
	return (*C.SDL_Renderer)(unsafe.Pointer(r))
}

func CreateRenderer(window *Window, index int, flags uint32) (*Renderer, error) {
	renderer := C.SDL_CreateRenderer(window.cptr(), C.int(index), C.Uint32(flags))
	if renderer == nil {
		return nil, GetError()
	}
	return (*Renderer)(unsafe.Pointer(renderer)), nil
}

func (renderer *Renderer) Destroy() error {
	lastErr := GetError()
	ClearError()
	C.SDL_DestroyRenderer(renderer.cptr())
	err := GetError()
	if err != nil {
		return err
	}
	SetError(lastErr)
	return nil
}

func (renderer *Renderer) Clear() error {
	return errorFromInt(int(
		C.SDL_RenderClear(renderer.cptr())))
}

func (renderer *Renderer) Copy(texture *Texture, src, dst *Rect) error {
	return errorFromInt(int(
		C.SDL_RenderCopy(
			renderer.cptr(),
			texture.cptr(),
			(*C.SDL_Rect)(unsafe.Pointer(src)),
			(*C.SDL_Rect)(unsafe.Pointer(dst)))))
}

func (renderer *Renderer) Present() {
	C.SDL_RenderPresent(renderer.cptr())
}

func (Renderer *Renderer) SetDrawColor(r, g, b, a int) {
	C.SDL_SetRenderDrawColor(Renderer.cptr(), C.Uint8(r), C.Uint8(g), C.Uint8(b), C.Uint8(a))
}
