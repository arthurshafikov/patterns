# Decorator Solution

To solve this problem, we'll need to create special classes called decorators (which look like Russian Matryoshka dolls, because they wrap each other) and the _Notifier_ interface which will describe only one method - _Send(msg string) error_. All the notifiers and decorators must implement this interface.

For example, we want to send a notification via phone + facebook:
1. We create **PhoneNotifier** instance
2. We pass this instance to **FacebookDecorator** constructor
3. We execute the job.

The program will call successively the decorator's _Send()_ method, which fires the facebook notification, and then will proceed to the phone notifier's _Send()_ method, which sends the notification via phone.
