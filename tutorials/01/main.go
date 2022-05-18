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

	fmt_ctx := avformat.AVFormatContext{}
	fmt_ctx.OpenInput(os.Args[1])
	defer fmt_ctx.CloseInput()

	fmt_ctx.FindStreamInfo()

	var codec *avcodec.AVCodec
	var avctx *avcodec.AVCodecContext
	for stream := range fmt_ctx.AVStreams() {
		if stream.CodecType() == avutil.AVMEDIA_TYPE_VIDEO {
			codec, _ = avcodec.FindDecoder(stream.CodecID())
			avctx, _ = avcodec.AllocContext3(codec)
			avctx.ParametersToContext(stream.CodecParameters())
			fmt.Println(avctx)
		}
	}
	avctx.Open2(codec)

}
