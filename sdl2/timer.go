package sdl2

/*
#include <SDL.h>
*/
import "C"

func Delay(ms uint32) {
	C.SDL_Delay(C.Uint32(ms))
}
