package remotes

import (
	"github.com/arthurshafikov/patterns/bridge/solution/devices"
)

type RemoteControl struct {
	device devices.Device // bridge.
}

func NewRemoteControl(device devices.Device) *RemoteControl {
	return &RemoteControl{
		device: device,
	}
}

func (r *RemoteControl) TogglePower() {
	if r.device.IsEnabled() {
		r.device.Disable()
	} else {
		r.device.Enable()
	}
}

func (r *RemoteControl) VolumeDown() {
	r.device.SetVolume(r.device.GetVolume() - 1)
}

func (r *RemoteControl) VolumeUp() {
	r.device.SetVolume(r.device.GetVolume() + 1)
}

func (r *RemoteControl) ChannelDown() {
	r.device.SetChannel(r.device.GetChannel() - 1)
}

func (r *RemoteControl) ChannelUp() {
	r.device.SetChannel(r.device.GetChannel() + 1)
}

func (r *RemoteControl) ShowInfo() {
	r.device.ShowInfo()
}
