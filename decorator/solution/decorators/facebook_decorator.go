package decorators

import "github.com/arthurshafikov/patterns/decorator/solution/notifiers"

type FacebookDecorator struct {
	notifier notifiers.Notifier
}

func NewFacebookDecorator(notifier notifiers.Notifier) *FacebookDecorator {
	return &FacebookDecorator{
		notifier: notifier,
	}
}

func (n *FacebookDecorator) Send(message string) error {
	if err := notifiers.NewFacebook().Send(message); err != nil {
		return err
	}

	return n.notifier.Send(message)
}
