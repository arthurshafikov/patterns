package notifiers

type Notifier interface {
	Send(message string) error
}
