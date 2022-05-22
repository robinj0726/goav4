package sdl2

/*
#include <SDL.h>
*/
import "C"
import "unsafe"

type Rect struct {
	X int32 // the x location of the rectangle's upper left corner
	Y int32 // the y location of the rectangle's upper left corner
	W int32 // the width of the rectangle
	H int32 // the height of the rectangle
}

func (a *Rect) cptr() *C.SDL_Rect {
	return (*C.SDL_Rect)(unsafe.Pointer(a))
}
