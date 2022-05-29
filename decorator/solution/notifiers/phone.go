package notifiers

import "fmt"

type Phone struct {
}

func NewPhone() *Phone {
	return &Phone{}
}

func (p *Phone) Send(msg string) error {
	fmt.Printf("Message \"%s\" was sent by Phone\n", msg)

	return nil
}
