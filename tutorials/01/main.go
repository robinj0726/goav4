package main

import (
	"fmt"

	"github.com/robinj730/goav4/avcodec"
	"github.com/robinj730/goav4/avformat"
	"github.com/robinj730/goav4/avutil"
)

func main() {
	// if len(os.Args) < 2 {
	// 	fmt.Println("Please provide a movie file")
	// 	return
	// }

	infile := "../sample.mp4"

	pFormatCtx := avformat.AVFormatContext{}
	pFormatCtx.OpenInput(infile)
	defer pFormatCtx.CloseInput()

	pFormatCtx.FindStreamInfo()

	var pCodec *avcodec.AVCodec
	var pCodecCtx *avcodec.AVCodecContext
	for stream := range pFormatCtx.AVStreams() {
		if stream.CodecType() == avutil.AVMEDIA_TYPE_VIDEO {
			pCodec, _ = avcodec.FindDecoder(stream.CodecID())
			pCodecCtx, _ = avcodec.AllocContext3(pCodec)
			pCodecCtx.ParametersToContext(stream.CodecParameters())
		}
	}
	err := pCodecCtx.Open2(pCodec)
	if err != nil {
		panic(err)
	}

	pFrame := avutil.FrameAlloc()
	defer pFrame.Free()

	pFrameRGB := avutil.FrameAlloc()
	defer pFrameRGB.Free()

	numBytes := avutil.GetImageBufferSize(int(avutil.AV_PIX_FMT_RGB24), pCodecCtx.Width(), pCodecCtx.Height(), 16)
	buffer := avutil.Malloc(numBytes)
	defer avutil.Free(buffer)

	avutil.FillImageArrays(pFrameRGB, buffer, int(avutil.AV_PIX_FMT_RGB24), pCodecCtx.Width(), pCodecCtx.Height(), 16)

	// fmt.Println(pCodecCtx)
	// sws_ctx, _ := swscale.GetContext(pCodecCtx.Width(), pCodecCtx.Height(), pCodecCtx.PixFmt(), pCodecCtx.Width(), pCodecCtx.Height(), int(avutil.AV_PIX_FMT_RGB24), swscale.SWS_BILINEAR)

	pkt := avcodec.PacketAlloc()
	defer pkt.Free()

	for {
		err := pFormatCtx.ReadFrame(pkt.PacketRef())
		if err != nil {
			break
		}

		fmt.Println(pkt)

	}
}
