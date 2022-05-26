package devices

type TV struct {
	enabled bool
	volume  int
	channel int
}

func NewTV() *TV {
	return &TV{}
}

func (tv TV) IsEnabled() bool {
	return tv.enabled
}

func (tv *TV) Enable() {
	tv.enabled = true
}

func (tv *TV) Disable() {
	tv.enabled = false
}

func (tv TV) GetVolume() int {
	return tv.volume
}

func (tv *TV) SetVolume(volume int) {
	tv.volume = volume
}

func (tv TV) GetChannel() int {
	return tv.channel
}

func (tv *TV) SetChannel(channel int) {
	tv.channel = channel
}
