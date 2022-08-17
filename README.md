# Go Programming (Golang): The Complete Developer's Guide by Udemy Course
This repo contains my exercises and annotations that I made during the course. Be comfortable to clone or talk to me about it :)

# Section 1: Introduction about the course and resources
You can find all of the course projects, code, slides at the below repository. You can use this as a guide as you go through the course, learn Go, and build projects!

https://github.com/jayson-lennon/ztm-golang

Additionally, this course utilizes exercise files which are packaged in a zip file attached to this lecture. The zip file contains an "src" folder which can be used throughout the entire course and a "solutions" folder which contains solutions to all projects and exercises.

Also, the repository for the exercises can be cloned from GitHub here:

git clone https://github.com/jayson-lennon/ztm-golang

To view solutions using the cloned repo, check out the solutions branch using: `git checkout solutions`


# Section 2: Introduction to Go (Golang)
- Packages are Go's way of organizing code;
	- Should be a single program or a single-purpose library
	- Can be published to and imported from the Go package registry
- Modules are a collection of packages
	- go.mod
- Data types is a way to specify how data should be interpreted;
	- Go uses static typing, which is checked at compile time
	- Converting between types requires parentheses;
- Go CLI Tool provides the go command line utility (go build, go run, go test, etc.)

# Section 3: Go programming Fundamentals
- Variables are a way to access memory locations using an alias
	- Multiple ways to create variables: single, compound, block, create e assign
	- Variables can be assigned to other variables
	- Variables names can only be used once per scope
	- Variables declared, but not assigned to, will have a default value
	- "Comma, ok" idiom allows you to reuse the second variable
- The main function is first function that we run in a program

- Functions
    - Encapsulate program functionality which leads to more maintainable code
    - Have parameters which define the input data
    - Are used by calling the function and supplying arguments
    - Can return multiple values
    - And underscore (_) can be used to ignore a return value