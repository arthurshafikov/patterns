package decorators

import "github.com/arthurshafikov/patterns/decorator/solution/notifiers"

type PhoneDecorator struct {
	notifier notifiers.Notifier
}

func NewPhoneDecorator(notifier notifiers.Notifier) *PhoneDecorator {
	return &PhoneDecorator{
		notifier: notifier,
	}
}

func (n *PhoneDecorator) Send(message string) error {
	if err := notifiers.NewPhone().Send(message); err != nil {
		return err
	}

	return n.notifier.Send(message)
}
