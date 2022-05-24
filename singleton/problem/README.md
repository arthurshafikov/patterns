# Singleton Problem

The problem is that we don't want to create multiple **Queue** objects, because there must be only one instance. We need to restrict creating multiple objects.

One of the possible solutions is to create the queue instance in _main()_ function and pass it to all functions, but imagine if there is a lot of places where **Queue** constructor is called. In that case, this solution seems very hard to implement.
