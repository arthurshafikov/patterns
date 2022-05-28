# Decorator Problem

Imagine we have an application and we have to set up a notification system in case something goes wrong. We have facebook, phone and slack notifiers. But what if we want to receive notification about major crashes through multiple notifiers? For example slack + phone or facebook + slack or even facebook + phone + slack. In order to do that we have to create a bunch of subclasses (**SlackNotifier**, **PhoneSlackNotifier**, **FacebookPhoneSlackNotifier** etc.).

Our code has bloated with redunant classes, but what if we add one more notifier - **EmailNotifier**, we'll need to create 7 subclasses - **EmailSlackNotifier**, **EmailPhoneNotifier**, **EmailFacebookNotifier**, **EmailPhoneSlackNotifier**, **EmailFacebookSlackNotifier**, **EmailFacebookPhoneNotifier**, **EmailFacebookPhoneSlackNotifier**. In my opinion, this solution is unacceptable and terrible.
