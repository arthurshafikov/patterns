package remotes

import (
	"fmt"

	"github.com/arthurshafikov/patterns/bridge/problem/devices"
)

type RemoteTVControl struct {
	tv *devices.TV
}

func NewRemoteTVControl() *RemoteTVControl {
	return &RemoteTVControl{
		tv: devices.NewTV(),
	}
}

func (r *RemoteTVControl) TogglePower() {
	if r.tv.IsEnabled() {
		r.tv.Disable()
	} else {
		r.tv.Enable()
	}
}

func (r *RemoteTVControl) VolumeDown() {
	r.tv.SetVolume(r.tv.GetVolume() - 1)
}

func (r *RemoteTVControl) VolumeUp() {
	r.tv.SetVolume(r.tv.GetVolume() + 1)
}

func (r *RemoteTVControl) ChannelDown() {
	r.tv.SetChannel(r.tv.GetChannel() - 1)
}

func (r *RemoteTVControl) ChannelUp() {
	r.tv.SetChannel(r.tv.GetChannel() + 1)
}

func (r *RemoteRadioControl) ShowInfo() {
	fmt.Printf("Radio volume: %v, radio channel: %v\n", r.radio.GetVolume(), r.radio.GetChannel())
}
