# Facade Solution

We created a new facade class - **VideoConverter**, which incapsulates all the library logic and initialization. In our client's code we don't know anything about this video library, we use it through our facade class. And if we are to update our code to a new version of the library, we can easily do it within the facade class.
