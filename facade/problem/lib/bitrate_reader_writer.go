package lib

type BitrateReaderWriter struct {
}

func NewBitrateReaderWriter() *BitrateReaderWriter {
	return &BitrateReaderWriter{}
}

func (br *BitrateReaderWriter) Read(filename string, codec any) []byte {
	// read file in given codec, return bytes

	return []byte{'s', 'o', 'm', 'e'}
}

func (br *BitrateReaderWriter) Convert(buffer []byte, codec any, outputFilename string) *VideoFile {
	// convert given buffer in file, using a given codec

	return NewVideoFile(outputFilename)
}

func (br *BitrateReaderWriter) Save(videoFile *VideoFile) bool {
	// save videoFile

	return true
}
