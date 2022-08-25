# Go Programming (Golang): The Complete Developer's Guide by Udemy Course
This repo contains my exercises and annotations that I made during the course. Be comfortable to clone or talk to me about it :)

## Summary

    * [Section 1: Introduction about the course and resources] (#section 1)
    * [Section 2: Introduction to Go (Golang)]
    * [Section 3: Go programming: Fundamentals] ()
    * [Section 4: Go programming: Types] ()
    * [Section 5: Idiomatic Go] ()
    * [Section 6: Interfaces in Go] ()
    * [Section 7: Concurrent Programming with Go] ()
    * [Section 8: Final Milestone Project: Pixl] ()
    * [Section 9: Project: MailingList Microservice] ()
    * [Section 10: Where to Go from here?] ()
    * [Section 11: BONUS SECTION] ()
## Section 1: Introduction about the course and resources
You can find all of the course projects, code, slides at the below repository. You can use this as a guide as you go through the course, learn Go, and build projects!

https://github.com/jayson-lennon/ztm-golang

Additionally, this course utilizes exercise files which are packaged in a zip file attached to this lecture. The zip file contains an "src" folder which can be used throughout the entire course and a "solutions" folder which contains solutions to all projects and exercises.

Also, the repository for the exercises can be cloned from GitHub here:

git clone https://github.com/jayson-lennon/ztm-golang

To view solutions using the cloned repo, check out the solutions branch using: `git checkout solutions`


## Section 2: Introduction to Go (Golang)
- Packages are Go's way of organizing code;
	- Should be a single program or a single-purpose library
	- Can be published to and imported from the Go package registry
- Modules are a collection of packages
	- go.mod
- Data types is a way to specify how data should be interpreted;
	- Go uses static typing, which is checked at compile time
	- Converting between types requires parentheses;
- Go CLI Tool provides the go command line utility (go build, go run, go test, etc.)

## Section 3: Go programming Fundamentals
# Variables 
    - Are a way to access memory locations using an alias
	- Multiple ways to create variables: single, compound, block, create e assign
	- Variables can be assigned to other variables
	- Variables names can only be used once per scope
	- Variables declared, but not assigned to, will have a default value
	- "Comma, ok" idiom allows you to reuse the second variable
- The main function is first function that we run in a program

# Functions
    - Encapsulate program functionality which leads to more maintainable code
    - Have parameters which define the input data
    - Are used by calling the function and supplying arguments
    - Can return multiple values
    - And underscore (_) can be used to ignore a return value

# Operators
    - Arithmetic Operators:
        + Addition          +=
        - Subtraction       -=
        * Multiplication    *=
        / Divisor           /=
        % Remainder         %=
    
        - Arithmetic Operators
            var total = 3 + 3
        - Arithmetic Assignment
            a := 1      a := 1
            a += 3      a = a + 3
        - Increment/Decrement
            i++
            i--
    ** Arithmetic operators always return a number **

    - Relational Operators
        <   Less Than
        <=  Less Than or Equal To
        >   Greater Than
        >=  Greater Than or Equal To
        ==  Equal To
        !=  Not Equal To

    - Logic Operators (true or false)
        &&  And
        ||  Or
        !   Not

# If... Else
    - If block is for true condition
    - Else block is for the false condition

    - Early returns should be used when possible for efficiency and code clarity
        token, err := getSession("rafaella")
        if err != nil {
            return
        }

# Switch
    - Can be use to check a variable for different values;
        - Use comma to check multiple values on a single case;
    - The 'fallthrough'keyword will execute the next case.

# Looping
    - The 'for' keyword creates a loop;
    - Use the 'break'keyword to exit the loop on a specific condition;
    - The initialization variable can be used only with the loop block;
        for i := 1; i < 10; i++; {}  // for i <10 {}  //  for {}
# Section Review - Dice Roller
    - In this class I'm gonna make a coding exercise, that will creating a program with some topics that we've covered in this section, such as:
        ➡ [Standard Library Packages](https://pkg.go.dev/std);
        ➡ [Generating Random Numbers] (https://pkg.go.dev/math/rand@go1.19);
        ➡ Control Flow;
        ➡ Looping;
        ➡ Retrieving System Time
    

