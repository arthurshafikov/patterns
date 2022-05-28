package job

import "github.com/arthurshafikov/patterns/decorator/problem/notifiers"

type Job struct {
	notifier notifiers.Notifier
	message  string
}

func NewJob() *Job {
	return &Job{}
}

func (j *Job) SetNotifier(notifier notifiers.Notifier) *Job {
	j.notifier = notifier

	return j
}

func (j *Job) SetMessage(message string) *Job {
	j.message = message

	return j
}

func (j *Job) Execute() error {
	return j.notifier.Send(j.message)
}
