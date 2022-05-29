package notifiers

import "fmt"

type Slack struct {
}

func NewSlack() *Slack {
	return &Slack{}
}

func (p *Slack) Send(msg string) error {
	fmt.Printf("Message \"%s\" was sent by Slack\n", msg)

	return nil
}
