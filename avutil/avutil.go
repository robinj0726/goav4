package avutil

/*
#cgo pkg-config: libavutil
#include <libavutil/avutil.h>
*/
import "C"

const (
	AVMEDIA_TYPE_UNKNOWN    = C.AVMEDIA_TYPE_UNKNOWN ///< Usually treated as AVMEDIA_TYPE_DATA
	AVMEDIA_TYPE_VIDEO      = C.AVMEDIA_TYPE_VIDEO
	AVMEDIA_TYPE_AUDIO      = C.AVMEDIA_TYPE_AUDIO
	AVMEDIA_TYPE_DATA       = C.AVMEDIA_TYPE_DATA ///< Opaque data information usually continuous
	AVMEDIA_TYPE_SUBTITLE   = C.AVMEDIA_TYPE_SUBTITLE
	AVMEDIA_TYPE_ATTACHMENT = C.AVMEDIA_TYPE_ATTACHMENT ///< Opaque data information usually sparse
	AVMEDIA_TYPE_NB         = C.AVMEDIA_TYPE_NB
)
