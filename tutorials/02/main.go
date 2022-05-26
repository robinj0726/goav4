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
	"time"

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

	pFormatCtx.DumpFormat()

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

	n := 0
	for {
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

			// C.SaveFrameToYUV((*C.struct_AVFrame)(pFrame.FrameRef()), (C.int)(pCodecCtx.Width()), (C.int)(pCodecCtx.Height()), (C.int)(n))

			sws_ctx.Scale(pFrame, 0, pCodecCtx.Height(), pFrameYUV)

			texture.UpdateYUV(nil, (*byte)(pFrame.DataPtr(0)), pFrame.LineSize(0), (*byte)(pFrame.DataPtr(1)), pFrame.LineSize(0), (*byte)(pFrame.DataPtr(2)), pFrame.LineSize(2))

			err = render.Clear()
			if err != nil {
				panic(err)
			}

			err = render.Copy(texture, nil, nil)
			if err != nil {
				panic(err)
			}

			render.Present()

			time.Sleep(30 * time.Millisecond)
			n += 1
		}
	}
}
