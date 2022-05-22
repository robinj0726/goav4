package sdl2

/*
#include <SDL.h>

static inline int RenderCopy(SDL_Renderer *renderer, SDL_Texture *texture, SDL_Rect *src, int dst_x, int dst_y, int dst_w, int dst_h)
{
	SDL_Rect dst = {dst_x, dst_y, dst_w, dst_h};
	return SDL_RenderCopy(renderer, texture, src, &dst);
}
*/
import "C"
import "unsafe"

const (
	TEXTUREACCESS_STATIC    = C.SDL_TEXTUREACCESS_STATIC    // changes rarely, not lockable
	TEXTUREACCESS_STREAMING = C.SDL_TEXTUREACCESS_STREAMING // changes frequently, lockable
	TEXTUREACCESS_TARGET    = C.SDL_TEXTUREACCESS_TARGET    // can be used as a render target
)

type Renderer C.SDL_Renderer

type Texture C.SDL_Texture

func (r *Renderer) cptr() *C.SDL_Renderer {
	return (*C.SDL_Renderer)(unsafe.Pointer(r))
}

func (t *Texture) cptr() *C.SDL_Texture {
	return (*C.SDL_Texture)(unsafe.Pointer(t))
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

func (renderer *Renderer) CreateTexture(format uint32, access int, w, h int32) (*Texture, error) {
	texture := C.SDL_CreateTexture(
		renderer.cptr(),
		C.Uint32(format),
		C.int(access),
		C.int(w),
		C.int(h))
	if texture == nil {
		return nil, GetError()
	}
	return (*Texture)(unsafe.Pointer(texture)), nil
}

func (texture *Texture) Destroy() error {
	lastErr := GetError()
	ClearError()
	C.SDL_DestroyTexture(texture.cptr())
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
	if dst == nil {
		return errorFromInt(int(
			C.SDL_RenderCopy(
				renderer.cptr(),
				texture.cptr(),
				src.cptr(),
				dst.cptr())))
	}
	return errorFromInt(int(
		C.RenderCopy(
			renderer.cptr(),
			texture.cptr(),
			src.cptr(),
			C.int(dst.X), C.int(dst.Y), C.int(dst.W), C.int(dst.H))))
}

func (renderer *Renderer) Present() {
	C.SDL_RenderPresent(renderer.cptr())
}

func (texture *Texture) UpdateYUV(rect *Rect, yPlane []byte, yPitch int, uPlane []byte, uPitch int, vPlane []byte, vPitch int) error {
	var yPlanePtr, uPlanePtr, vPlanePtr *byte
	if yPlane != nil {
		yPlanePtr = &yPlane[0]
	}
	if uPlane != nil {
		uPlanePtr = &uPlane[0]
	}
	if vPlane != nil {
		vPlanePtr = &vPlane[0]
	}
	return errorFromInt(int(
		C.SDL_UpdateYUVTexture(
			texture.cptr(),
			rect.cptr(),
			(*C.Uint8)(unsafe.Pointer(yPlanePtr)),
			C.int(yPitch),
			(*C.Uint8)(unsafe.Pointer(uPlanePtr)),
			C.int(uPitch),
			(*C.Uint8)(unsafe.Pointer(vPlanePtr)),
			C.int(vPitch))))
}
