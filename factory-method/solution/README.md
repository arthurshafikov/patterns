# Factory-Method Solution

The solution is in creating logistic classes which decide what instance of transport they should create. All transports' methods have been changed to implement _Transport_ interface. So our main programm doesn't care which transport it is, code can just call method that are defined in the interface. Also it doesn't matter which logistic type is needed in the programm, because all of them have an interface as well, which forces them to implement **CreateTransport()** method.

Therefore, now you can easily add new transport/logistic type and not to be concerned about new errors.

Pros:
- **SRP (Single Responsibility Principle)**, every class has only one reason to be changed.
- **OCP (Open/Closed Principle)**, it's easy to add any new deliveryType to the programm.

Cons:
- In this case code has become more complicated.
