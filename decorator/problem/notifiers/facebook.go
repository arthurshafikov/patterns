package notifiers

import "fmt"

type Facebook struct {
}

func NewFacebook() *Facebook {
	return &Facebook{}
}

func (p *Facebook) Send(msg string) error {
	fmt.Printf("Message \"%s\" was sent by Facebook\n", msg)

	return nil
}
