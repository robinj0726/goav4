package main

import (
	"fmt"
	"os"

	"github.com/robinj730/goav4/avcodec"
	"github.com/robinj730/goav4/avformat"
	"github.com/robinj730/goav4/avutil"
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

}
