package main

import (
	"fmt"
	"os"

	"github.com/robinj730/goav4/avcodec"
	"github.com/robinj730/goav4/avformat"
	"github.com/robinj730/goav4/avutil"
	"github.com/robinj730/goav4/avutil/avimage"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a movie file")
		return
	}

	pFormatCtx := avformat.AVFormatContext{}
	pFormatCtx.OpenInput(os.Args[1])
	defer pFormatCtx.CloseInput()

	pFormatCtx.FindStreamInfo()

	var pCodec *avcodec.AVCodec
	var pCodecCtx *avcodec.AVCodecContext
	for stream := range pFormatCtx.AVStreams() {
		if stream.CodecType() == avutil.AVMEDIA_TYPE_VIDEO {
			pCodec, _ = avcodec.FindDecoder(stream.CodecID())
			pCodecCtx, _ = avcodec.AllocContext3(pCodec)
			pCodecCtx.ParametersToContext(stream.CodecParameters())
			fmt.Println(pCodecCtx)
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

	numBytes := avimage.GetBufferSize(int(avutil.AV_PIX_FMT_RGB24), pCodecCtx.Width(), pCodecCtx.Height(), 16)
	buffer := avutil.Malloc(numBytes)
	defer avutil.Free(buffer)
}
