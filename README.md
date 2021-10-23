# go-command-pattern
Go command desing pattern example

D - Depency Invertion

This pattern is very useful to comply with this SOLID principle.

It precisely decouples the logic of "knowing" a sequence of actions by the client, since it ends up being encapsulated in a "commander" that has the sequence of actions and in turn the "commander" are abstracted behind an "invoker".

The "invoker" will be initialized with a specific "server" (object to manipulate) and will be able to execute various actions on it using the necessary "commanders".

The "invoker" exposes to the client the methods of the possible actions to be carried out on the manipulated object.

The sequence of actions of the "commander" could be implemented with a "chain of responsability" design pattern, but for simplicity it was not implemented in this project.

What is known as "receiver" in this pattern, is represented by the different types of servers that are in the source code and on which we can perform actions.

The "receivers" are manipulated by the "commanders"

<img src=https://upload.wikimedia.org/wikipedia/commons/c/c8/W3sDesign_Command_Design_Pattern_UML.jpg />
