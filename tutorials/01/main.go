package main

import (
	"fmt"
	"os"

	"github.com/robinj730/goav4/avformat"
	"github.com/robinj730/goav4/avutil"
)

func main() {
	fmt_ctx := avformat.AVFormatContext{}
	fmt_ctx.OpenInput(os.Args[1])
	defer fmt_ctx.CloseInput()

	fmt_ctx.FindStreamInfo()

	for codecpar := range fmt_ctx.AVStreams() {
		fmt.Printf("%#v\n", codecpar)
		if codecpar.CodecType == avutil.AVMEDIA_TYPE_VIDEO {

		}
	}

}
