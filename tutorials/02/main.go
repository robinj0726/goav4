package main

/*
#include <libavutil/frame.h>

static void SaveFrameToYUV(AVFrame *pFrame, int w, int h, int iFrame) {
  FILE *pFile;
  char szFilename[32];
  int height[3] = {h, h/2, h/2};

  // Open file
  sprintf(szFilename, "frame%d.yuv", iFrame);
  pFile=fopen(szFilename, "wb");
  if(pFile==NULL)
    return;

  // Write pixel data
  for (int i=0; i<3; i++) {
	for(int y=0; y<height[i]; y++) {
		fwrite(pFrame->data[i]+y*pFrame->linesize[i], 1, pFrame->linesize[i], pFile);
	}
  }

  // Close file
  fclose(pFile);
}
*/
import "C"
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

	err := sdl2.Init(sdl2.INIT_VIDEO)
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

	avutil.ImageAlloc(pFrameYUV, pCodecCtx.Width(), pCodecCtx.Height(), pCodecCtx.PixFmt(), 8)

	// fmt.Println(pCodecCtx)
	sws_ctx, _ := swscale.GetContext(pCodecCtx.Width(), pCodecCtx.Height(), pCodecCtx.PixFmt(), pCodecCtx.Width(), pCodecCtx.Height(), int(avutil.AV_PIX_FMT_YUV420P), swscale.SWS_BILINEAR)

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

	yPlaneSz := pCodecCtx.Width() * pCodecCtx.Height()
	yPlane := *(*[]uint8)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(pFrameYUV.Plane(0)),
		Len:  int(yPlaneSz),
		Cap:  int(yPlaneSz),
	}))

	uvPlaneSz := yPlaneSz / 4
	uPlane := *(*[]uint8)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(pFrameYUV.Plane(1)),
		Len:  int(uvPlaneSz),
		Cap:  int(uvPlaneSz),
	}))

	vPlane := *(*[]uint8)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(pFrameYUV.Plane(2)),
		Len:  int(uvPlaneSz),
		Cap:  int(uvPlaneSz),
	}))

	yPitch := pCodecCtx.Width()
	uvPitch := yPitch / 2

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

			// C.SaveFrameToYUV((*C.struct_AVFrame)(pFrame.FrameRef()), (C.int)(pCodecCtx.Width()), (C.int)(pCodecCtx.Height()), (C.int)(n))

			sws_ctx.Scale(pFrame, 0, pCodecCtx.Height(), pFrameYUV)
			texture.UpdateYUV(nil, yPlane, int(yPitch), uPlane, int(uvPitch), vPlane, int(uvPitch))

			render.Clear()
			render.Copy(texture, nil, nil)
			render.Present()

			n += 1
		}
	}
}
