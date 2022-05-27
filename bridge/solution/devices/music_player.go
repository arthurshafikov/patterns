package devices

import "fmt"

type MusicPlayer struct {
	enabled bool
	volume  int
	channel int // playlist.
}

func NewMusicPlayer() *MusicPlayer {
	return &MusicPlayer{}
}

func (r MusicPlayer) IsEnabled() bool {
	return r.enabled
}

func (r *MusicPlayer) Enable() {
	r.enabled = true
}

func (r *MusicPlayer) Disable() {
	r.enabled = false
}

func (r MusicPlayer) GetVolume() int {
	return r.volume
}

func (r *MusicPlayer) SetVolume(volume int) {
	r.volume = volume
}

func (r MusicPlayer) GetChannel() int {
	return r.channel
}

func (r *MusicPlayer) SetChannel(channel int) {
	r.channel = channel
}

func (r MusicPlayer) ShowInfo() {
	fmt.Printf("MusicPlayer volume: %v, musicPlayer channel: %v\n", r.volume, r.channel)
}
