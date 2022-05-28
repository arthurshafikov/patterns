package notifiers

import "fmt"

type FacebookSlack struct {
}

func NewFacebookSlack() *FacebookSlack {
	return &FacebookSlack{}
}

func (p *FacebookSlack) Send(msg string) error {
	fmt.Printf("Message \"%s\" was sent by Facebook + Slack\n", msg)

	return nil
}
