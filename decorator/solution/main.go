package main

import (
	"fmt"

	"github.com/arthurshafikov/patterns/decorator/solution/decorators"
	"github.com/arthurshafikov/patterns/decorator/solution/job"
	"github.com/arthurshafikov/patterns/decorator/solution/notifiers"
)

func main() {
	job := job.NewJob().SetMessage("production is stopped working...")

	SingleOnes(job)
	CompoundOnes(job)
	/*
		Message "production is stopped working..." was sent by Phone
		Message "production is stopped working..." was sent by Facebook
		Message "production is stopped working..." was sent by Slack
		=======================
		Message "production is stopped working..." was sent by Slack
		Message "production is stopped working..." was sent by Phone
		=======================
		Message "production is stopped working..." was sent by Facebook
		Message "production is stopped working..." was sent by Phone
		=======================
		Message "production is stopped working..." was sent by Slack
		Message "production is stopped working..." was sent by Facebook
		=======================
		Message "production is stopped working..." was sent by Slack
		Message "production is stopped working..." was sent by Phone
		Message "production is stopped working..." was sent by Facebook
	*/
}

// have not changed.
func SingleOnes(job *job.Job) {
	var notifier notifiers.Notifier

	notifier = notifiers.NewPhone()
	job.SetNotifier(notifier).Execute()

	notifier = notifiers.NewFacebook()
	job.SetNotifier(notifier).Execute()

	notifier = notifiers.NewSlack()
	job.SetNotifier(notifier).Execute()
}

func CompoundOnes(job *job.Job) {
	var notifier notifiers.Notifier

	// Phone + Slack
	fmt.Println("=======================")
	notifier = notifiers.NewPhone()
	notifier = decorators.NewSlackDecorator(notifier)
	job.SetNotifier(notifier).Execute()

	// Phone + Facebook
	fmt.Println("=======================")
	notifier = notifiers.NewPhone()
	notifier = decorators.NewFacebookDecorator(notifier)
	job.SetNotifier(notifier).Execute()

	// Facebook + Slack
	fmt.Println("=======================")
	notifier = notifiers.NewFacebook()
	notifier = decorators.NewSlackDecorator(notifier)
	job.SetNotifier(notifier).Execute()

	// Facebook + Phone + Slack
	fmt.Println("=======================")
	notifier = notifiers.NewFacebook()
	notifier = decorators.NewPhoneDecorator(notifier)
	notifier = decorators.NewSlackDecorator(notifier)
	job.SetNotifier(notifier).Execute()
}
