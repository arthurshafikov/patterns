# Facade Problem

The problem is that we have an external library which converts files from one format to another. In order to use that library we need to write a huge initialization and make sure to do it right. Imagine if our program was incredibly big. What it would take to update our program to a new version of this video converter library? Or worse, to change that library to any other, that would be a disaster. Our code depends on this library.
