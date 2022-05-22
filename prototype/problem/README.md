# Prototype Problem

Imagine that we have an array of different **Shapes** - **Circles** and **Rectangles**, and suddenly we find out that we need to create an exact copy of every **Shape** in that array. And while we can copy some public properties, we cannot reach private ones, we could use struct copying, but that would copy our _uniqueID_ property, and we definitely don't want that. Also we make a dependency in our code on concrete classes (structures) of **Shapes** (**Circle** and **Rectangle**).

These are 2 main problems and that's where the **Prototype** pattern could be applied.
