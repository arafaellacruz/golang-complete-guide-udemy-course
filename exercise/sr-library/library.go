//--Summary:
//  Create a program to manage lending of library books.
//
//--Requirements:
//* The library must have books and members, and must include:
//  - Which books have been checked out
//  - What time the books were checked out
//  - What time the books were returnel
//* Perform the following:
//  - Add at least 4 books and at least 3 members to the library
//  - Check out a book
//  - Check in a book
//  - Print out initial library information, and after each change
//* There must only ever be one copy of the library in memory at any time
//
//--Notes:
//* Use the `time` package from the standard library for check in/out times
//* Liberal use of type aliases, structs, and maps will help organize this project

package main

import (
	"fmt"
	"time"
)

type Title string
type Name string
type LendAudit struct {
	checkIn  time.Time
	checkOut time.Time
}
type Member struct {
	name  Name
	books map[Title]LendAudit
}

type BookEntry struct {
	total  int
	lended int
}

type Library struct {
	members map[Name]Member
	books   map[Title]BookEntry
}

func printMemberAudit(member *Member) {
	for title, audit := range member.books {
		var returnTime string
		if audit.checkIn.IsZero() {
			returnTime = "[not returned yet]"
		} else {
			returnTime = audit.checkIn.String()
		}
		fmt.Println(member.name, ":", title, ":", audit.checkOut.String(), "through", returnTime)
	}
}

func printMemberAudits(library *Library) {
	for _, member := range library.members {
		printMemberAudit(&member)
	}
}

func printLibraryBooks(library *Library) {
	fmt.Println()
	for title, book := range library.books {
		fmt.Println(title, "/ total:", book.total, "/ lended:", book.lended)
	}
	fmt.Println()
}

func checkOutBook(library *Library, title Title, member *Member) bool {
	book, found := library.books[title]
	if !found {
		fmt.Println("Book not part of library")
		return false
	}
	if book.lended == book.total {
		fmt.Println("No more books available to lend")
		return false
	}
	book.lended += 1
	library.books[title] = book

	member.books[title] = LendAudit{checkOut: time.Now()}
	return true
}

func returnBook(library *Library, title Title, member *Member) bool {
	book, found := library.books[title]
	if !found {
		fmt.Println("Book not part of library")
		return false
	}

	audit, found := member.books[title]
	if !found {
		fmt.Println("Member did not check out this books")
		return false
	}

	book.lended -= 1
	library.books[title] = book

	audit.checkIn = time.Now()
	member.books[title] = audit
	return true
}

func main() {
	library := Library{
		books:   make(map[Title]BookEntry),
		members: make(map[Name]Member),
	}

	library.books["Harry Potter and the Philosopher's Stone"] = BookEntry{
		total:  4,
		lended: 0,
	}
	library.books["Harry Potter and the Chamber of Secrets"] = BookEntry{
		total:  3,
		lended: 0,
	}
	library.books["Harry Potter and the Prisoner of Azkaban"] = BookEntry{
		total:  2,
		lended: 0,
	}
	library.books["Harry Potter and the Goblet of Fire"] = BookEntry{
		total:  1,
		lended: 0,
	}

	library.members["Rafaella Cruz"] = Member{"Rafaella Cruz", make(map[Title]LendAudit)}
	library.members["Billie Eilish"] = Member{"Billie Eilish", make(map[Title]LendAudit)}
	library.members["Beyoncé Carter"] = Member{"Beyoncé Carter", make(map[Title]LendAudit)}

	fmt.Println("\nInitial:")
	printLibraryBooks(&library)
	printMemberAudits(&library)

	//  - Check out a book
	member := library.members["Rafaella Cruz"]
	checkedOut := checkOutBook(&library, "Harry Potter and the Philosopher's Stone", &member)
	fmt.Println("\nCheck out a book:")
	if checkedOut {
		printLibraryBooks(&library)
		printMemberAudits(&library)
	}

	fmt.Println("\nAfter lend some book:")
	printLibraryBooks(&library)
	printMemberAudits(&library)

	//  - Check in a book
	returned := returnBook(&library, "Harry Potter and the Philosopher's Stone", &member)
	fmt.Println("\nCheck in a book:")
	if returned {
		printLibraryBooks(&library)
		printMemberAudits(&library)
	}
}
