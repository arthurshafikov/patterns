# Prototype Solution

To solve this problem we need to create a common interface **Shape** for all shapes with _Clone()_ method in it. This method will allow to clone all of shape's properties and generate a uniqueID for each cloned shape. We got rid of dependency on concrete shape classes as well using this interface. So as a result we have a convenient way to copy any shape to use it in our programm.
