package avformat

/*
#cgo pkg-config: libavformat
#include <libavformat/avformat.h>
*/
import "C"

func AvformatVersion() uint {
	return uint(C.avformat_version())
}
