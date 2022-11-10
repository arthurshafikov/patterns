# Chain Of Responsibility Problem

The main problem is that we have to duplicate code that checks if user is authentificated/banned/an admin. The example here is pretty simple, but in real world, there would be much more duplicated code which would be very hard to change without errors and bugs. There also would be extra steps like brute-force protection/data encryption/request validation etc. We simply don't want to duplicate that code.
