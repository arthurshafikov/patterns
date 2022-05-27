package main

import (
	"github.com/arthurshafikov/patterns/bridge/solution/devices"
	"github.com/arthurshafikov/patterns/bridge/solution/remotes"
)

func main() {
	radio := devices.NewRadio()
	radioRemoteControl := remotes.NewRemoteControl(radio)

	radioRemoteControl.VolumeUp()
	radioRemoteControl.VolumeUp()
	radioRemoteControl.VolumeUp()
	radioRemoteControl.ChannelUp()
	radioRemoteControl.ChannelUp()
	radioRemoteControl.ShowInfo()
	// Radio volume: 3, radio channel: 2.

	tv := devices.NewTV()
	tvRemoteControl := remotes.NewRemoteControl(tv)

	tvRemoteControl.VolumeUp()
	tvRemoteControl.VolumeUp()
	tvRemoteControl.VolumeDown()
	tvRemoteControl.ChannelUp()
	tvRemoteControl.ChannelUp()
	tvRemoteControl.ChannelDown()
	tvRemoteControl.ShowInfo()
	// TV volume: 1, TV channel: 1.

	musicPlayer := devices.NewMusicPlayer()
	musicPlayerRemoteControl := remotes.NewRemoteControl(musicPlayer)

	musicPlayerRemoteControl.VolumeUp()
	musicPlayerRemoteControl.VolumeUp()
	musicPlayerRemoteControl.VolumeUp()
	musicPlayerRemoteControl.VolumeUp()
	musicPlayerRemoteControl.VolumeDown()
	musicPlayerRemoteControl.ChannelUp()
	musicPlayerRemoteControl.ChannelUp()
	musicPlayerRemoteControl.ChannelDown()
	musicPlayerRemoteControl.ShowInfo()
	// MusicPlayer volume: 3, musicPlayer channel: 1.
}
