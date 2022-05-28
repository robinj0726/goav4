package main

// typedef unsigned char Uint8;
// void AudioCallback(void *userdata, Uint8 *stream, int len);
import "C"
import (
	"fmt"
	"reflect"
	"unsafe"

	"github.com/robinj730/goav4/avcodec"
	"github.com/robinj730/goav4/avformat"
	"github.com/robinj730/goav4/avutil"
	"github.com/robinj730/goav4/sdl2"
)

var audioq chan *avcodec.AVPacket

func init() {
	audioq = make(chan *avcodec.AVPacket, 10)
}

func main() {
	infile := "../sample.mp4"

	err := sdl2.Init(sdl2.INIT_VIDEO | sdl2.INIT_AUDIO | sdl2.INIT_TIMER)
	if err != nil {
		panic(err)
	}

	pFormatCtx := avformat.AVFormatContext{}
	pFormatCtx.OpenInput(infile)
	defer pFormatCtx.CloseInput()

	pFormatCtx.FindStreamInfo()

	pFormatCtx.DumpFormat()

	var pCodec *avcodec.AVCodec
	var pCodecCtx *avcodec.AVCodecContext
	videoStream := -1

	var aCodec *avcodec.AVCodec
	var aCodecCtx *avcodec.AVCodecContext
	audioStream := -1

	i := 0
	for stream := range pFormatCtx.AVStreams() {
		if stream.CodecType() == avutil.AVMEDIA_TYPE_VIDEO {
			videoStream = i
			pCodec, _ = avcodec.FindDecoder(stream.CodecID())
			pCodecCtx, _ = avcodec.AllocContext3(pCodec)
			pCodecCtx.ParametersToContext(stream.CodecParameters())
		}

		if stream.CodecType() == avutil.AVMEDIA_TYPE_AUDIO {
			audioStream = i
			aCodec, _ = avcodec.FindDecoder(stream.CodecID())
			aCodecCtx, _ = avcodec.AllocContext3(aCodec)
			aCodecCtx.ParametersToContext(stream.CodecParameters())
		}
		i += 1
	}

	if videoStream == -1 {
		panic("no video stream inside")
	}

	if audioStream == -1 {
		panic("no audio stream inside")
	}

	wanted_spec := sdl2.AudioSpec{
		Freq:     int32(aCodecCtx.SampleRate()),
		Format:   sdl2.AUDIO_S16SYS,
		Channels: uint8(aCodecCtx.Channels()),
		Silence:  0,
		Samples:  1024,
		Callback: sdl2.AudioCallback(C.AudioCallback),
		UserData: aCodecCtx.CPtr(),
	}
	spec := sdl2.AudioSpec{}
	err = sdl2.OpenAudio(&wanted_spec, &spec)
	if err != nil {
		panic(err)
	}
	defer sdl2.CloseAudio()

	sdl2.PauseAudio(0)

	// open video codec
	err = pCodecCtx.Open2(pCodec)
	if err != nil {
		panic(err)
	}

	// open audio codec
	err = aCodecCtx.Open2(aCodec)
	if err != nil {
		panic(err)
	}

	pkt := avcodec.PacketAlloc()
	defer pkt.Free()

	for {
		pFrame := avutil.FrameAlloc()
		defer pFrame.Free()

		err := pFormatCtx.ReadFrame(pkt.PacketPtr())
		if err != nil {
			break
		}

		if pkt.StreamIndex() == videoStream {
			err := pCodecCtx.SendPacket(pkt.PacketPtr())
			if err != nil {
				panic(err)
			}

			err = pCodecCtx.ReceiveFrame(pFrame.FramePtr())
			if err != nil {
				panic(err)
			}
		}

		if pkt.StreamIndex() == audioStream {
			audioq <- pkt.PacketClone()
		}
	}
}

//export AudioCallback
func AudioCallback(userdata unsafe.Pointer, stream *C.Uint8, length C.int) {
	if stream == nil {
		return
	}

	n := int(length)
	dst := *(*[]C.Uint8)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(stream)),
		Len:  n,
		Cap:  n,
	}))

	aCodecCtx := avcodec.NewContextWithCPtr(userdata)
	pFrame := avutil.FrameAlloc()
	defer pFrame.Free()

	pkt := <-audioq
	defer pkt.Free()

	err := aCodecCtx.SendPacket(pkt.PacketPtr())
	if err != nil {
		panic(err)
	}

	err = aCodecCtx.ReceiveFrame(pFrame.FramePtr())
	if err != nil {
		panic(err)
	}

	// data_size := avutil.GetSamplesBufferSize(nil, aCodecCtx.Channels(),
	// 	pFrame.Samples(), aCodecCtx.SampleFmt(), 1)
	src := *(*[]C.Uint8)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(pFrame.DataPtr(0))),
		Len:  pFrame.LineSize(0),
		Cap:  pFrame.LineSize(0),
	}))

	fmt.Printf("%d, %d\n", len(dst), len(src))
	for i := 0; i < n; i++ {
		dst[i] = src[i]
	}
	sdl2.Delay(100)
}
