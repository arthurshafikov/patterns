# Adapter Problem

Imagine we are working for a SpaceX company, and we were given a task where we need to estimate hours and fuel that are needed on the flight to several planets. And usually, we receive the distance to these planets in kilometers, but we connect to some external library which gives us a distance only in miles. Therefore, we need to convert the miles to kilometers every time we need to estimate time and fuel for a flight to any planet. 

Time has passed. Everything seemed OK until a new programmer came who didn't know about our system and he was given a task to calculate the flight to Saturn. He used miles and passed it to the spaceship methods, which returned the wrong result and that's it, the ship won't have enought fuel to fly to the Saturn.
