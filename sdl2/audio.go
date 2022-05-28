package sdl2

/*
#include <SDL.h>
*/
import "C"
import "unsafe"

const (
	AUDIO_S8 = C.AUDIO_S8 // signed 8-bit samples
	AUDIO_U8 = C.AUDIO_U8 // unsigned 8-bit samples

	AUDIO_S16LSB = C.AUDIO_S16LSB // signed 16-bit samples in little-endian byte order
	AUDIO_S16MSB = C.AUDIO_S16MSB // signed 16-bit samples in big-endian byte order
	AUDIO_S16SYS = C.AUDIO_S16SYS // signed 16-bit samples in native byte order
	AUDIO_S16    = C.AUDIO_S16    // AUDIO_S16LSB
	AUDIO_U16LSB = C.AUDIO_U16LSB // unsigned 16-bit samples in little-endian byte order
	AUDIO_U16MSB = C.AUDIO_U16MSB // unsigned 16-bit samples in big-endian byte order
	AUDIO_U16SYS = C.AUDIO_U16SYS // unsigned 16-bit samples in native byte order
	AUDIO_U16    = C.AUDIO_U16    // AUDIO_U16LSB

	AUDIO_S32LSB = C.AUDIO_S32LSB // 32-bit integer samples in little-endian byte order
	AUDIO_S32MSB = C.AUDIO_S32MSB // 32-bit integer samples in big-endian byte order
	AUDIO_S32SYS = C.AUDIO_S32SYS // 32-bit integer samples in native byte order
	AUDIO_S32    = C.AUDIO_S32    // AUDIO_S32LSB

	AUDIO_F32LSB = C.AUDIO_F32LSB // 32-bit floating point samples in little-endian byte order
	AUDIO_F32MSB = C.AUDIO_F32MSB // 32-bit floating point samples in big-endian byte order
	AUDIO_F32SYS = C.AUDIO_F32SYS // 32-bit floating point samples in native byte order
	AUDIO_F32    = C.AUDIO_F32    // AUDIO_F32LSB
)

type AudioFormat uint16

type AudioCallback C.SDL_AudioCallback

type AudioSpec struct {
	Freq     int32          // DSP frequency (samples per second)
	Format   AudioFormat    // audio data format
	Channels uint8          // number of separate sound channels
	Silence  uint8          // audio buffer silence value (calculated)
	Samples  uint16         // audio buffer size in samples (power of 2)
	_        uint16         // padding
	Size     uint32         // audio buffer size in bytes (calculated)
	Callback AudioCallback  // the function to call when the audio device needs more data
	UserData unsafe.Pointer // a pointer that is passed to callback (otherwise ignored by SDL)
}

func OpenAudio(desired, obtained *AudioSpec) error {
	if C.SDL_OpenAudio(desired.cptr(), obtained.cptr()) != 0 {
		return GetError()
	}
	return nil
}

func CloseAudio() {
	C.SDL_CloseAudio()
}

func (as *AudioSpec) cptr() *C.SDL_AudioSpec {
	return (*C.SDL_AudioSpec)(unsafe.Pointer(as))
}

func PauseAudio(pauseOn int) {
	C.SDL_PauseAudio(C.int(pauseOn))
}
