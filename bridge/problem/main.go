package main

import "github.com/arthurshafikov/patterns/bridge/problem/remotes"

func main() {
	radioRemoteControl := remotes.NewRemoteRadioControl()

	radioRemoteControl.VolumeUp()
	radioRemoteControl.VolumeUp()
	radioRemoteControl.VolumeUp()

	radioRemoteControl.ChannelUp()
	radioRemoteControl.ChannelUp()

	radioRemoteControl.ShowInfo()
	// Radio volume: 3, radio channel: 2.

	tvRemoteControl := remotes.NewRemoteTVControl()

	tvRemoteControl.VolumeUp()
	tvRemoteControl.VolumeUp()
	tvRemoteControl.VolumeDown()

	tvRemoteControl.ChannelUp()
	tvRemoteControl.ChannelUp()
	tvRemoteControl.ChannelDown()

	tvRemoteControl.ShowInfo()
	// TV volume: 1, TV channel: 1.
}
