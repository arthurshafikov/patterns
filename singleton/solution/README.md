# Singleton Solution

To solve this problem, we created a new method _GetInstance()_ in **Queue** pkg, which returns an existing **Queue** instance or creates and returns a new one. Also, we made our **Queue** struct and its constructor private, therefore we restricted any **Queue** objects creation outside **Queue** pkg.

As a result, we have only one instance of **Queue** and that's exactly what the Singleton is.
