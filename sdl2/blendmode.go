package sdl2

/*
#include <SDL.h>
*/
import "C"
import "unsafe"

type BlendMode uint32

const (
	BLENDMODE_NONE    = C.SDL_BLENDMODE_NONE  // no blending
	BLENDMODE_BLEND   = C.SDL_BLENDMODE_BLEND // alpha blending
	BLENDMODE_ADD     = C.SDL_BLENDMODE_ADD   // additive blending
	BLENDMODE_MOD     = C.SDL_BLENDMODE_MOD   // color modulate
	BLENDMODE_MUL     = C.SDL_BLENDMODE_MUL   // color multiply
	BLENDMODE_INVALID = C.SDL_BLENDMODE_INVALID
)

func (bm BlendMode) c() C.SDL_BlendMode {
	return C.SDL_BlendMode(C.Uint32(bm))
}

func (bm *BlendMode) cptr() *C.SDL_BlendMode {
	return (*C.SDL_BlendMode)(unsafe.Pointer(bm))
}
