# Builder Solution

So, the solution is that we need to add additional classes called **builders**. In them we create convenient methods for each param that we used in the constructor of a **House** class. As a result, we have a **House** class with its own methods (for example _Sell()_ or _Demolish()_) and **House Builder** class (with _BuildGarage()_, _BuildRooms()_ etc.).

Thanks to that the code readability has improved and there is no more huge calls to constructor with redunant params.

## Directors

It's not required, but you can incapsulate the most used builder calls to the **Director** classes. If the structure of the house depends on the cost you pass, you can easily create a director and resolve what house you need to get. 

Also, it's not necessary to build only houses, in my example I've also added a **Hostel Builder**, and thanks to the _EstateBuilder_ interface, the director doesn't need to know which builder he is using, he just calls it's methods that were described in the interface.
