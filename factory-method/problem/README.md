# Factory-Method Problem

So basically, the problem is that the logic of creating specific types of transport is scattered around the application, and everywhere there is the switch/case block which determines what transport it should use. If we imagine that we need to add more types of delivery to our program, there would be a huge amount of code that we need to change and that could cause new errors in our working application.

As a result, our code is full of these switch/case blocks and it's inconvenient to add any new transport type (or change/remove the old one).
