package main

import (
	"fmt"

	"github.com/robinj730/goav4/avformat"
)

func main() {
	fmt.Println(avformat.Version())

	// avio_ctx_buffer_size := 4096
	fmt_ctx := avformat.AllocContext()
	defer fmt_ctx.FreeContext()

	// avio_ctx_buffer := avutil.AvMalloc(avio_ctx_buffer_size)

}
