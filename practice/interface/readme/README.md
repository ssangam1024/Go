
Struct:

A struct in Go is just a way to group different pieces of information together in one unit. It helps to organize data and make your code easier to manage.

Interface:

Interface defines set of rules/ methods that a type must follow.  a type like struct has all required methods (means the specific functions that are written inside the interface), it automatically implements that interface. 

* Struct = A box to store information (like Car, Bicycle).
* Interface = A set of rules for different types to follow, like saying "Any type that can drive should have a Drive method."
You use an interface when you want to group things that do similar actions but may do them in different ways.

Why Use an Interface?
* Simplifies code: You can write a function that works with any type that implements the interface. You don't need to write separate code for Car, Bicycle, and Truck—just one function that works with anything that can drive.
* Flexibility: You can easily add new types (like Motorcycle) without changing the code that uses the interface.


What is an Interface?
An interface in Go is like a set of rules or a list of actions that a type (like a struct) must be able to do. But you don't care about the exact type of the object—you just care that it can do the things in the list (the actions).
Think of it like a contract:
* If a type can follow the contract, then it is said to "implement" the interface.
* You don't need to know how it does those things, just that it can do them.


Simple Example:
Imagine you have different types of vehicles:
* Car
* Bicycle
* Truck
They all need to drive. You don't care whether it's a car or a bike—you just want to be able to tell them to drive.
So you create an interface called Drivable:
........
Now, Car, Bicycle, and Truck can all have their own version of the Drive method. As long as they all can drive, they are said to "implement" the Drivable interface.


