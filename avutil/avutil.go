package avutil

/*
#cgo pkg-config: libavutil
#include <libavutil/avutil.h>
*/
import "C"

type AVMediaType int

const (
	AVMEDIA_TYPE_UNKNOWN    AVMediaType = -1 ///< Usually treated as AVMEDIA_TYPE_DATA
	AVMEDIA_TYPE_VIDEO      AVMediaType = 0
	AVMEDIA_TYPE_AUDIO      AVMediaType = 1
	AVMEDIA_TYPE_DATA       AVMediaType = 2 ///< Opaque data information usually continuous
	AVMEDIA_TYPE_SUBTITLE   AVMediaType = 3
	AVMEDIA_TYPE_ATTACHMENT AVMediaType = 4 ///< Opaque data information usually sparse
	AVMEDIA_TYPE_NB         AVMediaType = 5
)
