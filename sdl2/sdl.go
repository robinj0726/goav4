package sdl2

/*
#cgo pkg-config: sdl2
#include <SDL.h>
*/
import "C"

const (
	INIT_TIMER          = C.SDL_INIT_TIMER          // timer subsystem
	INIT_AUDIO          = C.SDL_INIT_AUDIO          // audio subsystem
	INIT_VIDEO          = C.SDL_INIT_VIDEO          // video subsystem; automatically initializes the events subsystem
	INIT_JOYSTICK       = C.SDL_INIT_JOYSTICK       // joystick subsystem; automatically initializes the events subsystem
	INIT_HAPTIC         = C.SDL_INIT_HAPTIC         // haptic (force feedback) subsystem
	INIT_GAMECONTROLLER = C.SDL_INIT_GAMECONTROLLER // controller subsystem; automatically initializes the joystick subsystem
	INIT_EVENTS         = C.SDL_INIT_EVENTS         // events subsystem
	INIT_NOPARACHUTE    = C.SDL_INIT_NOPARACHUTE    // compatibility; this flag is ignored
	INIT_SENSOR         = C.SDL_INIT_SENSOR         // sensor subsystem
	INIT_EVERYTHING     = C.SDL_INIT_EVERYTHING     // all of the above subsystems
)

func Init(flags uint32) error {
	if C.SDL_Init(C.Uint32(flags)) != 0 {
		return GetError()
	}
	return nil
}

func Quit() {
	C.SDL_Quit()
}
