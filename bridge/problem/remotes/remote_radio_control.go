package remotes

import (
	"fmt"

	"github.com/arthurshafikov/patterns/bridge/problem/devices"
)

type RemoteRadioControl struct {
	radio *devices.Radio
}

func NewRemoteRadioControl() *RemoteRadioControl {
	return &RemoteRadioControl{
		radio: devices.NewRadio(),
	}
}

func (r *RemoteRadioControl) TogglePower() {
	if r.radio.IsEnabled() {
		r.radio.Disable()
	} else {
		r.radio.Enable()
	}
}

func (r *RemoteRadioControl) VolumeDown() {
	r.radio.SetVolume(r.radio.GetVolume() - 1)
}

func (r *RemoteRadioControl) VolumeUp() {
	r.radio.SetVolume(r.radio.GetVolume() + 1)
}

func (r *RemoteRadioControl) ChannelDown() {
	r.radio.SetChannel(r.radio.GetChannel() - 1)
}

func (r *RemoteRadioControl) ChannelUp() {
	r.radio.SetChannel(r.radio.GetChannel() + 1)
}

func (r *RemoteTVControl) ShowInfo() {
	fmt.Printf("TV volume: %v, TV channel: %v\n", r.tv.GetVolume(), r.tv.GetChannel())
}
