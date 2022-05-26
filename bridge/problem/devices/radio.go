package devices

type Radio struct {
	enabled bool
	volume  int
	channel int
}

func NewRadio() *Radio {
	return &Radio{}
}

func (r Radio) IsEnabled() bool {
	return r.enabled
}

func (r *Radio) Enable() {
	r.enabled = true
}

func (r *Radio) Disable() {
	r.enabled = false
}

func (r Radio) GetVolume() int {
	return r.volume
}

func (r *Radio) SetVolume(volume int) {
	r.volume = volume
}

func (r Radio) GetChannel() int {
	return r.channel
}

func (r *Radio) SetChannel(channel int) {
	r.channel = channel
}
