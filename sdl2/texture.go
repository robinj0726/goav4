package sdl2

/*
#include <SDL.h>
*/
import "C"
import (
	"reflect"
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
			rect.cptr(),
			(*C.Uint8)(unsafe.Pointer(yPlanePtr)),
			C.int(yPitch),
			(*C.Uint8)(unsafe.Pointer(uPlanePtr)),
			C.int(uPitch),
			(*C.Uint8)(unsafe.Pointer(vPlanePtr)),
			C.int(vPitch))))
}

func (texture *Texture) Lock(rect *Rect) ([]byte, int, error) {
	var _pitch C.int
	var _pixels unsafe.Pointer
	var b []byte
	var length int

	ret := C.SDL_LockTexture(texture.cptr(), rect.cptr(), &_pixels, &_pitch)
	if ret < 0 {
		return b, int(_pitch), GetError()
	}

	_, _, w, h, err := texture.Query()
	if err != nil {
		return b, int(_pitch), GetError()
	}

	pitch := int32(_pitch)
	if rect != nil {
		bytesPerPixel := pitch / w
		length = int(bytesPerPixel * (w*rect.H - rect.X - (w - rect.X - rect.W)))
	} else {
		length = int(pitch * h)
	}
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sliceHeader.Cap = int(length)
	sliceHeader.Len = int(length)
	sliceHeader.Data = uintptr(_pixels)

	return b, int(pitch), nil
}

func (texture *Texture) Unlock() {
	C.SDL_UnlockTexture(texture.cptr())
}

func (texture *Texture) Query() (format uint32, access int, width int32, height int32, err error) {
	var _format C.Uint32
	var _access C.int
	var _width C.int
	var _height C.int

	ret := int(C.SDL_QueryTexture(texture.cptr(), &_format, &_access, &_width, &_height))

	format = uint32(_format)
	access = int(_access)
	width = int32(_width)
	height = int32(_height)
	err = errorFromInt(ret)

	return
}
