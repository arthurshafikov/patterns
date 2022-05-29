package decorators

import "github.com/arthurshafikov/patterns/decorator/solution/notifiers"

type SlackDecorator struct {
	notifier notifiers.Notifier
}

func NewSlackDecorator(notifier notifiers.Notifier) *SlackDecorator {
	return &SlackDecorator{
		notifier: notifier,
	}
}

func (n *SlackDecorator) Send(message string) error {
	if err := notifiers.NewSlack().Send(message); err != nil {
		return err
	}

	return n.notifier.Send(message)
}
