package lib

type VideoFile struct {
	Filename string
}

func NewVideoFile(filename string) *VideoFile {
	return &VideoFile{
		Filename: filename,
	}
}
