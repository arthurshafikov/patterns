# Flyweight Solution

The most amount of memory of our program was taken by the texture field in the **Tree** class, but that field contained one of the 5 treeTypes, so, in every **Tree** instance, we duplicated this texture. To change that, we created a TreeFactory, which stores every possible TreeType under the hood and does not duplicate that data among the trees. Therefore, our **Tree** classes do not contain this heavy texture field, but have a link to a **TreeType**. 

Now our benchmark looks like this:
```
    Memory allocations : 10 007 438
    Memory bytes : 703 676 384
    Time taken: 2.270890492s
```

And, comparing to the old one:
```
    Memory allocations : 10 007 438 (10 003 738)
    Memory bytes : 703 676 384 (1 263 430 016) // almost halved
    Time taken: 2.270890492s
```

As a result, our program started to consume much less memory.

