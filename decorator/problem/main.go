package main

import (
	"fmt"

	"github.com/arthurshafikov/patterns/decorator/problem/job"
	"github.com/arthurshafikov/patterns/decorator/problem/notifiers"
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
		Message "production is stopped working..." was sent by Phone + Slack
		=======================
		Message "production is stopped working..." was sent by Facebook + Phone
		=======================
		Message "production is stopped working..." was sent by Facebook + Slack
		=======================
		Message "production is stopped working..." was sent by Facebook + Phone + Slack
	*/
}

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
	notifier = notifiers.NewPhoneSlack()
	job.SetNotifier(notifier).Execute()

	// Phone + Facebook
	fmt.Println("=======================")
	notifier = notifiers.NewFacebookPhone()
	job.SetNotifier(notifier).Execute()

	// Facebook + Slack
	fmt.Println("=======================")
	notifier = notifiers.NewFacebookSlack()
	job.SetNotifier(notifier).Execute()

	// Facebook + Phone + Slack
	fmt.Println("=======================")
	notifier = notifiers.NewFacebookPhoneSlack()
	job.SetNotifier(notifier).Execute()
}
