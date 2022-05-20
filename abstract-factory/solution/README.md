# Abstract-Factory Solution

- The first step to solve our problem was to create interfaces for each type of furniture. 
- Then we created a factory interface and proceeded to making factories for each furniture family we have.
- The third step was to use these new factories in our code.

So now our furniture are creating using these factories and not through constructors, what gives us an opportunity to easily create furniture families together without fear to mix them up.

Also, to make our code even more strict, we could make all individual furniture consturctor functions unavailable everywhere but factories. In that case, we could be sure that nobody will create an instance of furniture individually.
