# Automatas-Builder
This repository contains a project used to create and test differents automatas.

You have to compile and run the main.go file to use the program.

You can download and install Go from the official page: https://go.dev

In the folder named "files" there is a file with the name "(0,1)-(0)(0).graph" this is an example of an automata that
accept any string that ends with "01" using the Alphabet E = {0, 1}. The name of the file is like the regular
expression -> [0-1]*[0][1], but the file name really doesn't matter.

This program is not complete and some features are not yet implemented.

## How to use
Before this you have to download the project and run it using ``` go run main.go ```

### [1]. Add a new node.
If you want to create a node (state) remember that the name of a node is unique, so you can't have two nodes with same names.
When a node is created this doesn'e have any connection.

### [2]. Modify a node.
This is just to change the name of a node.

### [3]. Create connection.
This is used to create a connection between two nodes by an arc, this arc can use more that one letter, for example:
If you want to create a connection like q0 -- 0, 1 --> q1 when then program ask for the characther transition you can write "01" and the connection will be created like that.

### [4]. Remove a connection.
This is just to remove a connection.

### [5]. See the transitions table.
Here you can see the delta (transitions) array.

> [!NOTE]
> This documentation is not completed.
