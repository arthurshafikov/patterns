package lib

type AudioMixer struct {
}

func NewAudioMixer() *AudioMixer {
	return &AudioMixer{}
}

func (a *AudioMixer) FixAudioInVideoFile(file *VideoFile) *VideoFile {
	// fix audio...

	return file
}
