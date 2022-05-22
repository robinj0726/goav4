package sdl2

/*
#include <SDL.h>
#include "events.h"
*/
import "C"
import "unsafe"

const (
	// Application events
	QUIT = C.SDL_QUIT // user-requested quit
)

type Event interface {
	GetType() uint32      // GetType returns the event type
	GetTimestamp() uint32 // GetTimestamp returns the timestamp of the event
}

type CEvent struct {
	Type uint32
	_    [52]byte // padding
}

type CommonEvent struct {
	Type      uint32 // the event type
	Timestamp uint32 // timestamp of the event
}

func (e *CommonEvent) GetType() uint32 {
	return e.Type
}

func (e *CommonEvent) GetTimestamp() uint32 {
	return e.Timestamp
}

type QuitEvent struct {
	Type      uint32 // QUIT
	Timestamp uint32 // timestamp of the event
}

func (e *QuitEvent) GetType() uint32 {
	return e.Type
}

func (e *QuitEvent) GetTimestamp() uint32 {
	return e.Timestamp
}

func PollEvent() Event {
	ret := C.PollEvent()
	if ret == 0 {
		return nil
	}
	return goEvent((*CEvent)(unsafe.Pointer(&C.event)))
}

func goEvent(cevent *CEvent) Event {
	switch cevent.Type {
	case QUIT:
		return (*QuitEvent)(unsafe.Pointer(cevent))
	default:
		return (*CommonEvent)(unsafe.Pointer(cevent))
	}
}
