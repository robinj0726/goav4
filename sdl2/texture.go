package sdl2

/*
#include <SDL.h>
*/
import "C"
import (
	"unsafe"
)

type Texture C.SDL_Texture

func (t *Texture) cptr() *C.SDL_Texture {
	return (*C.SDL_Texture)(unsafe.Pointer(t))
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

func (texture *Texture) UpdateYUV(rect *Rect, yPlanePtr *byte, yPitch int, uPlanePtr *byte, uPitch int, vPlanePtr *byte, vPitch int) error {
	return errorFromInt(int(
		C.SDL_UpdateYUVTexture(
			texture.cptr(),
			(*C.SDL_Rect)(unsafe.Pointer(rect)),
			(*C.Uint8)(unsafe.Pointer(yPlanePtr)),
			C.int(yPitch),
			(*C.Uint8)(unsafe.Pointer(uPlanePtr)),
			C.int(uPitch),
			(*C.Uint8)(unsafe.Pointer(vPlanePtr)),
			C.int(vPitch))))
}

func (texture *Texture) SetBlendMode(bm BlendMode) error {
	return errorFromInt(int(
		C.SDL_SetTextureBlendMode(texture.cptr(), bm.c())))
}
