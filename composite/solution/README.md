# Composite Solution

We need to unify products and boxes under a common interface - **HasPrice**, and all we need to do is to call this method on the box and it'll return the price of all products in it.

The greatest benefit of this approach is that you don’t need to care about the concrete classes of objects that compose the tree. You don’t need to know whether an object is a simple product or a sophisticated box.  When you call a method, the objects themselves pass the request down the tree.
