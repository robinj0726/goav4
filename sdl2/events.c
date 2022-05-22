#include "_cgo_export.h"
#include "events.h"

SDL_Event event;

int PollEvent()
{
	return SDL_PollEvent(&event);
}