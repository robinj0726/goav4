package sdl2

/*
#include <SDL.h>
*/
import "C"

type Rect struct {
	X int32 // the x location of the rectangle's upper left corner
	Y int32 // the y location of the rectangle's upper left corner
	W int32 // the width of the rectangle
	H int32 // the height of the rectangle
}
