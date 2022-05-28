package notifiers

import "fmt"

type FacebookPhoneSlack struct {
}

func NewFacebookPhoneSlack() *FacebookPhoneSlack {
	return &FacebookPhoneSlack{}
}

func (p *FacebookPhoneSlack) Send(msg string) error {
	fmt.Printf("Message \"%s\" was sent by Facebook + Phone + Slack\n", msg)

	return nil
}
