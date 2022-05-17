package main

import (
	"fmt"
	"os"

	"github.com/robinj730/goav4/avcodec"
	"github.com/robinj730/goav4/avformat"
	"github.com/robinj730/goav4/avutil"
)

func main() {
	if len(os.Args) <= 2 {
		fmt.Printf("Usage: %s <input file> <output file>\n", os.Args[0])
		return
	}

	infile := os.Args[1]
	// outfile := os.Args[2]

	pkt := avcodec.PacketAlloc()
	defer pkt.Free()

	fmt_ctx := avformat.AVFormatContext{}
	fmt_ctx.OpenInput(infile)
	defer fmt_ctx.CloseInput()

	fmt_ctx.FindStreamInfo()

	var codec *avcodec.AVCodec
	var avctx *avcodec.AVCodecContext
	for codecpar := range fmt_ctx.AVStreams() {
		if codecpar.CodecType == avutil.AVMEDIA_TYPE_VIDEO {
			codec, _ = avcodec.FindDecoder(codecpar.CodecID)
			avctx, _ = avcodec.AllocContext3(codec)

			codec.AVCodecParameters = codecpar
			break
		}
	}

	if codec == nil {
		panic("Codec not found")
	}

	if avctx == nil {
		panic("Could not allocate video codec context")
	}

	_, err := avcodec.ParserInit(codec.CodecID)
	if err != nil {
		panic(err)
	}

	err = avctx.Open2(codec)
	if err != nil {
		panic("Could not open codec")
	}

	frame := avutil.FrameAlloc()
	defer frame.Free()
}
