package avcodec

import "github.com/robinj730/goav4/avutil"

type AVCodecParameters struct {
	CodecType avutil.AVMediaType
	CodecID   int
}
