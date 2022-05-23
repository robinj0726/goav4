package main

import (
	"reflect"
	"unsafe"

	"github.com/robinj730/goav4/avcodec"
	"github.com/robinj730/goav4/avformat"
	"github.com/robinj730/goav4/avutil"
	"github.com/robinj730/goav4/sdl2"
	"github.com/robinj730/goav4/swscale"
)

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

	var pCodec *avcodec.AVCodec
	var pCodecCtx *avcodec.AVCodecContext
	videoStream := -1
	i := 0
	for stream := range pFormatCtx.AVStreams() {
		if stream.CodecType() == avutil.AVMEDIA_TYPE_VIDEO {
			pCodec, _ = avcodec.FindDecoder(stream.CodecID())
			pCodecCtx, _ = avcodec.AllocContext3(pCodec)
			pCodecCtx.ParametersToContext(stream.CodecParameters())
			videoStream = i
			break
		}
		i += 1
	}

	if videoStream == -1 {
		panic("no video stream inside")
	}

	err = pCodecCtx.Open2(pCodec)
	if err != nil {
		panic(err)
	}

	pFrame := avutil.FrameAlloc()
	defer pFrame.Free()

	pFrameYUV := avutil.FrameAlloc()
	defer pFrameYUV.Free()

	avutil.ImageAlloc(pFrameYUV, pCodecCtx.Width(), pCodecCtx.Height(), pCodecCtx.PixFmt(), 16)

	// fmt.Println(pCodecCtx)
	sws_ctx, _ := swscale.GetContext(pCodecCtx.Width(), pCodecCtx.Height(), pCodecCtx.PixFmt(), pCodecCtx.Width(), pCodecCtx.Height(), int(avutil.AV_PIX_FMT_RGB24), swscale.SWS_BILINEAR)

	pkt := avcodec.PacketAlloc()
	defer pkt.Free()

	window, err := sdl2.CreateWindow("ffmpeg tutorial 02", sdl2.WINDOWPOS_UNDEFINED, sdl2.WINDOWPOS_UNDEFINED,
		pCodecCtx.Width(), pCodecCtx.Height(), sdl2.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	render, err := sdl2.CreateRenderer(window, -1, 0)
	if err != nil {
		panic(err)
	}
	defer render.Destroy()

	texture, err := render.CreateTexture(sdl2.PIXELFORMAT_YV12, sdl2.TEXTUREACCESS_STREAMING, pCodecCtx.Width(), pCodecCtx.Height())
	if err != nil {
		panic(err)
	}
	defer texture.Destroy()

	yPlane := *(*[]uint8)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(pFrameYUV.Plane(0)),
		Len:  pFrameYUV.PlaneSize(0),
		Cap:  pFrameYUV.PlaneSize(0),
	}))

	uPlane := *(*[]uint8)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(pFrameYUV.Plane(1)),
		Len:  pFrameYUV.PlaneSize(1),
		Cap:  pFrameYUV.PlaneSize(1),
	}))

	vPlane := *(*[]uint8)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(pFrameYUV.Plane(2)),
		Len:  pFrameYUV.PlaneSize(2),
		Cap:  pFrameYUV.PlaneSize(2),
	}))

	n := 0
	for {
		err := pFormatCtx.ReadFrame(pkt.PacketRef())
		if err != nil {
			break
		}

		if pkt.StreamIndex() == videoStream {
			err := pCodecCtx.SendPacket(pkt.PacketRef())
			if err != nil {
				panic(err)
			}

			err = pCodecCtx.ReceiveFrame(pFrame.FrameRef())
			if err != nil {
				panic(err)
			}

			sws_ctx.Scale(*pFrame, 0, pCodecCtx.Height(), *pFrameYUV)
			texture.UpdateYUV(nil, yPlane, len(yPlane), uPlane, len(uPlane), vPlane, len(vPlane))

			render.Clear()
			render.Copy(texture, nil, nil)
			render.Present()
			n += 1
		}
	}
}