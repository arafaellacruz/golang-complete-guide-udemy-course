# Go Programming (Golang)
This repo contains my exercises and annotations that I made during the course "Go Programming (Golang): The Complete Developer's Guide by Udemy Course". Be comfortable to clone or talk to me about it :)

## Summary

- [Section 1: Introduction about the course and resources](https://github.com/rafacruzz/golang-complete-guide-udemy-course#section-1-introduction-about-the-course-and-resources)
- [Section 2: Introduction to Go (Golang)](https://github.com/rafacruzz/golang-complete-guide-udemy-course#section-1-introduction-about-the-course-and-resources)
- [Section 3: Go programming: Fundamentals](https://github.com/rafacruzz/golang-complete-guide-udemy-course#section-1-introduction-about-the-course-and-resources)
- [Section 4: Go programming: Types](https://github.com/rafacruzz/golang-complete-guide-udemy-course#section-4-types)
- [Section 5: Idiomatic Go](https://github.com/rafacruzz/golang-complete-guide-udemy-course#section-5-idiomatic-go)
- [Section 6: Interfaces in Go](https://github.com/rafacruzz/golang-complete-guide-udemy-course#section-6-interfaces-in-go)
- [Section 7: Concurrent Programming with Go](https://github.com/rafacruzz/golang-complete-guide-udemy-course#section-7-concurrent-programming-with-go)
- [Section 8: Final Milestone Project: Pixl](https://github.com/rafacruzz/golang-complete-guide-udemy-course#section-8-final-milestone-project:-pixl)
- [Section 9: Project: MailingList Microservice](https://github.com/rafacruzz/golang-complete-guide-udemy-course#section-9-project-mailinglist-microservice)
- [Section 10: Where to Go from here?](https://github.com/rafacruzz/golang-complete-guide-udemy-course#section-10-where-to-go-from-here?)
- [Section 11: BONUS SECTION](https://github.com/rafacruzz/golang-complete-guide-udemy-course#section-11-bonus-section)
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
### Variables 
- Are a way to access memory locations using an alias
- Multiple ways to create variables: single, compound, block, create e assign
- Variables can be assigned to other variables
- Variables names can only be used once per scope
- Variables declared, but not assigned to, will have a default value
- "Comma, ok" idiom allows you to reuse the second variable

*The main function is first function that we run in a program*

### Functions
- Encapsulate program functionality which leads to more maintainable code
- Have parameters which define the input data
- Are used by calling the function and supplying arguments
- Can return multiple values
- And underscore (_) can be used to ignore a return value

### Operators
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

*Arithmetic operators always return a number*

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

### If... Else
- If block is for true condition
- Else block is for the false condition

- Early returns should be used when possible for efficiency and code clarity
    token, err := getSession("rafaella")
    if err != nil {
        return
    }

### Switch
- Can be use to check a variable for different values;
    - Use comma to check multiple values on a single case;
- The 'fallthrough'keyword will execute the next case.

### Looping
- The 'for' keyword creates a loop;
- Use the 'break'keyword to exit the loop on a specific condition;
- The initialization variable can be used only with the loop block;
    for i := 1; i < 10; i++; {}  // for i <10 {}  //  for {}
### Section Review - Dice Roller
- In this class I'm gonna make a coding exercise, that will creating a program with some topics that we've covered in this section, such as:
        <br>
        ➡ [Standard Library Packages](https://pkg.go.dev/std);
        <br>
        ➡ [Generating Random Numbers](https://pkg.go.dev/math/rand@go1.19);
        <br>
        ➡ Control Flow;
        <br>
        ➡ Looping;
        <br>
        ➡ Retrieving System Time.
    
## Section 4: Types

### Structures
- Use to group similar data;
- Are similar to a "class" in other programming languages;
- Defining a Structure: We'll use the keyword 'type' followed by the name of the structure and then finally the 'struct' keyword;
- Default Values: Any fields not indicated during instantiation will have default values;
- Accessing Fields: fields can be read from and written to.

### Arrays
- Are fixed-size collections of same-type items;
- Are accessed using an *array index*;
- Array elements can be optionally set during array creation;
- Elements not manually assigned a value will have a default;
- Use the *len()* function to iterate arrays in a *for* loop;
- Whenever we're working with the iterator variable itself within a for loop, *always make a copy*;

### Slices
- Slices are a more convenient way to work with arrays;
```
    Example: nameOfSlice := []typeOfSlice{"itemsOfTheSlice, x, y, z"}
```
- Can be resized using the *append()* function;
- Can be created with a special *slice syntax*;
- Slice memory can be preallocated using the *make()* function;
- Slices always require an underlying array.
- The make built-in function allocates and initializes an object of type slice, map, or chan (only). Like new, the first argument is a type, not a value. Unlike new, make's return type is the same as the type of its argument, not a pointer to it.

#### Ranges
- The *range* keyword create an iterator for us, automatically. It's kind of like an easier way of making a 'for'loop with the counter.
```
slice := []string{"Hello", "world", "!"}
// i always start with the 0 index.
// And when we use 'range'we're going to get two things back from a slice, the index and the and the content of the index.
for i, element := range slice {
    fmt.Println(i, element)
}
```

### Maps
- Maps store data in key-value pairs;
- Very fast key accesses;
- Use *range* to iterate through a map;
- Use the *make()* function to create an empty map;
- Use the *delete()* function to remove an entry from the map;
- Read and write map uses similar syntax to array.
- [Example from DEMO](https://github.com/rafacruzz/golang-complete-guide-udemy-course/blob/main/exercise/maps/maps.demo.go)

### Pointers
- We use the (*) whe you're declaring that you need a pointer to something and we use (&) when we want to actually create a pointer from something that already exists; 
- Example:
```
value := 10
var valuePtr  *int
valuePtr = &value
```
- Pointers are used to modify data that exists outside of a function;
- Asterisk (*) on a type indicates the type is a pointer;
- Ampersand (&) creates a pointer;
- Asterisk (*) on a variable will deference the pointer;
    - Operations on a dereferenced pointer occur on the original data.

### Section Review - Library

## Section 5: Idiomatic Go

### Receiver Functions
- Receiver functions provide the "dot" notation for structs;
    - Create a more convenient API's;
- *Pointer* receivers can modify a struct;
- *Value* receivers can not modify a struct;
- Common to use pointer receivers.

### iota

- The *iota* keyword can be used to assign integers to constants;
- *const* is like a variable, but unchanging;
- Common to make groups of constants;
- *iota* keyword can be used to automatically assign values;
    - Increments 1 in the next constance;
- Long-form and short-form;
- Skip values with '_';
- *iota* values can be expressions (iota + 5);
- Use a receiver function to more easily work with constants and *iota*.

### Text formattinh: fmt
- *fmt* package provides terminal printing and string formatting;
- Provides functions:
    - Printf: custom format;
    - Print: simple print;
    - Println: simple print with a newline;
- *F* and *S* variants of the above functions:
    - *F* prints to a data stream: *Fprintf*, *Fprint*, *Fprintln*;
    - *S* prints to a new string: *Sprintf*, *Sprint*, *Sprintln*;

### Packages

### Init function
- Common to have an initialization step in programs;
    - Creaating a function named *init()* will perform initialization;
    - *init()* is ran before the main() function;
- Allows creation and validation of program state before execution begins;
    - Check network connection, database connections, cache expensive operations, etc;
- Each package can have it's own *init()* function;
- All packages will execute *init()* before *main()* runs.

### Testing
- Test files end with *_test*;
- Unit testing: test individual functions;
- Integration testing: test functions/modules working together;
- Unit tests should be in the same package;
- Execute *go test* to run your tests;
- Available Testing Functions:
    - Many testing functions available in the *testing* package;
    - *Fail()* - Mark the test as failed;
        - *Errorf(string)* - Fail & add a message;
    - *FailNow()* - Mark the test as failed, abort current test;
        - *Fatal(string)* - Fail, abort, and add a message;
    - *Logf()* - Equivalent to *Printf*, but only when test fails

- Test Tables
    - Usually need to test more than one set of input data (can be used to test multiple pieces of data);
    - *Test tables* can be used to accomplish this;
        - Similar to parameterized testing;


## Section 6: Interfaces in Go