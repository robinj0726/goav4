package sdl2

/*
#include <SDL.h>

void GoSetError(const char *fmt) {
  SDL_SetError("%s", fmt);
}
*/
import "C"
import "errors"

var emptyCString *C.char = C.CString("")

func GetError() error {
	if err := C.SDL_GetError(); err != nil {
		gostr := C.GoString(err)
		// SDL_GetError returns "an empty string if there hasn't been an error message"
		if len(gostr) > 0 {
			return errors.New(gostr)
		}
	}
	return nil
}

func ClearError() {
	C.SDL_ClearError()
}

func SetError(err error) {
	if err != nil {
		C.GoSetError(C.CString(err.Error()))
		return
	}
	C.GoSetError(emptyCString)
}

func errorFromInt(code int) (err error) {
	if code < 0 {
		err = GetError()
		if err == nil {
			err = errors.New("Unknown error (probably using old version of SDL2 and the function called is not supported?)")
		}
	}
	return
}
