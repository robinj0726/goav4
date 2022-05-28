package main

/*
#cgo pkg-config: libavutil

#include <libavutil/frame.h>
#include <SDL2/SDL_rect.h>

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

static void calculate_display_rect(SDL_Rect *rect,
                                   int scr_xleft, int scr_ytop, int scr_width, int scr_height,
                                   int pic_width, int pic_height)
{
    AVRational aspect_ratio = {1,1};
    int64_t width, height, x, y;

    if (av_cmp_q(aspect_ratio, av_make_q(0, 1)) <= 0)
        aspect_ratio = av_make_q(1, 1);

    aspect_ratio = av_mul_q(aspect_ratio, av_make_q(pic_width, pic_height));

    height = scr_height;
    width = av_rescale(height, aspect_ratio.num, aspect_ratio.den) & ~1;
    if (width > scr_width) {
        width = scr_width;
        height = av_rescale(width, aspect_ratio.den, aspect_ratio.num) & ~1;
    }
    x = (scr_width - width) / 2;
    y = (scr_height - height) / 2;
    rect->x = scr_xleft + x;
    rect->y = scr_ytop  + y;
    rect->w = FFMAX((int)width,  1);
    rect->h = FFMAX((int)height, 1);
}

*/
import "C"
import (
	"unsafe"

	"github.com/robinj730/goav4/avcodec"
	"github.com/robinj730/goav4/avformat"
	"github.com/robinj730/goav4/avutil"
	"github.com/robinj730/goav4/sdl2"
)

func main() {
	infile := "../sample.mp4"

	err := sdl2.Init(sdl2.INIT_VIDEO | sdl2.INIT_TIMER)
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

	// pFrameYUV := avutil.FrameAlloc()
	// defer pFrameYUV.Free()

	// fmt.Println(pCodecCtx)
	// sws_ctx, _ := swscale.GetContext(pCodecCtx.Width(), pCodecCtx.Height(), pCodecCtx.PixFmt(), pCodecCtx.Width(), pCodecCtx.Height(), int(avutil.AV_PIX_FMT_YUV420P), swscale.SWS_BILINEAR)

	pkt := avcodec.PacketAlloc()
	defer pkt.Free()

	window, err := sdl2.CreateWindow("ffmpeg tutorial 02", sdl2.WINDOWPOS_UNDEFINED, sdl2.WINDOWPOS_UNDEFINED,
		pCodecCtx.Width(), pCodecCtx.Height(), sdl2.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	render, err := sdl2.CreateRenderer(window, -1, sdl2.RENDERER_ACCELERATED|sdl2.RENDERER_PRESENTVSYNC)
	if err != nil {
		panic(err)
	}
	defer render.Destroy()

	texture, err := render.CreateTexture(sdl2.PIXELFORMAT_IYUV, sdl2.TEXTUREACCESS_STREAMING, pCodecCtx.Width(), pCodecCtx.Height())
	if err != nil {
		panic(err)
	}
	defer texture.Destroy()
	texture.SetBlendMode(sdl2.BLENDMODE_NONE)

	n := 0
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

			// pFrameYUV := avutil.FrameAlloc()
			// avutil.ImageAlloc(pFrameYUV, pCodecCtx.Width(), pCodecCtx.Height(), pCodecCtx.PixFmt(), 16)

			// sws_ctx.Scale(pFrame, 0, pCodecCtx.Height(), pFrameYUV)

			render.SetDrawColor(0, 0, 0, 255)
			err = render.Clear()
			if err != nil {
				panic(err)
			}

			// C.SaveFrameToYUV((*C.struct_AVFrame)(pFrame.FramePtr()),
			// 	(C.int)(pCodecCtx.Width()), (C.int)(pCodecCtx.Height()), (C.int)(n))

			texture.UpdateYUV(nil, (*byte)(pFrame.DataPtr(0)), pFrame.LineSize(0),
				(*byte)(pFrame.DataPtr(1)), pFrame.LineSize(1),
				(*byte)(pFrame.DataPtr(2)), pFrame.LineSize(2))

			rect := sdl2.Rect{}
			C.calculate_display_rect((*C.SDL_Rect)(unsafe.Pointer(&rect)), 0, 0, C.int(pCodecCtx.Width()), C.int(pCodecCtx.Height()),
				C.int(pCodecCtx.Width()), C.int(pCodecCtx.Height()))

			err = render.CopyEx(texture, nil, nil, 0, nil, sdl2.FLIP_NONE)
			if err != nil {
				panic(err)
			}

			render.Present()

			sdl2.Delay(60)
			n += 1
		}
	}
}
