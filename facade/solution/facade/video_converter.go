package facade

import (
	"github.com/arthurshafikov/patterns/facade/solution/lib"
)

type VideoConverter struct {
}

func NewVideoConverter() *VideoConverter {
	return &VideoConverter{}
}

func (vc *VideoConverter) Convert(filename, outputFilename, format string) bool {
	file := lib.NewVideoFile(filename)
	sourceCodec := lib.NewCodecFactory().Extract(file)

	var destinationCodec any
	switch format {
	case "ogg":
		destinationCodec = lib.NewOggCompressionCodec()
	case "mp4":
		destinationCodec = lib.NewMPEG4CompressionCodec()
	}

	bitrateReaderWriter := lib.NewBitrateReaderWriter()

	buffer := bitrateReaderWriter.Read(filename, sourceCodec)
	outputFile := bitrateReaderWriter.Convert(buffer, destinationCodec, outputFilename)

	audioMixer := lib.NewAudioMixer()
	outputFile = audioMixer.FixAudioInVideoFile(outputFile)

	return bitrateReaderWriter.Save(outputFile)
}
