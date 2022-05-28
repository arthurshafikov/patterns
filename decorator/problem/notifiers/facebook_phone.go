package notifiers

import "fmt"

type FacebookPhone struct {
}

func NewFacebookPhone() *FacebookPhone {
	return &FacebookPhone{}
}

func (p *FacebookPhone) Send(msg string) error {
	fmt.Printf("Message \"%s\" was sent by Facebook + Phone\n", msg)

	return nil
}
