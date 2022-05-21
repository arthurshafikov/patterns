# Builder Problem

Let's say that there are a lot of variations how you can build the **House** (_hasGarage_, _hasGarden_, etc.). In order to implement that, we need to either create a bunch of subclasses (**HouseWithGarage**, **HouseWithGarden**) or put a lot of arguments in **House** constructor.

Thus, calls to this constructor become huge and most of the passing params are redunant. And it's not obvious what it means when you pass params like this (_ModernStyle_, _3_, _false_, _false_, _true_, _false_), and this leads to worse code readability and maintainability.
