package lib

import (
	"strings"
)

type CodecFactory struct {
}

func NewCodecFactory() *CodecFactory {
	return &CodecFactory{}
}

func (c *CodecFactory) Extract(file *VideoFile) any {
	splittedFilename := strings.Split(file.Filename, ".")

	if len(splittedFilename) < 2 {
		panic("wrong filename")
	}

	switch splittedFilename[1] {
	case "ogg":
		return NewOggCompressionCodec()
	case "mp4":
		return NewMPEG4CompressionCodec()
	}

	panic("unknown format " + splittedFilename[1])
}
