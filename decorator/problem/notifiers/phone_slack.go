package notifiers

import "fmt"

type PhoneSlack struct {
}

func NewPhoneSlack() *PhoneSlack {
	return &PhoneSlack{}
}

func (p *PhoneSlack) Send(msg string) error {
	fmt.Printf("Message \"%s\" was sent by Phone + Slack\n", msg)

	return nil
}
