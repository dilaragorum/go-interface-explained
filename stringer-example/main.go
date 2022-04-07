package main

import (
	"fmt"
	"log"
	"strconv"
)

// Declare a Book type which satisfies the fmt.Stringer interface(String() string)
type Book struct {
	Title  string
	Author string
}

func (b Book) String() string {
	return fmt.Sprintf("Book: %s - %s", b.Title, b.Author)
}

//Declare a Count type which satisfies the fmt.Stringer interface(String() string)
type Count int

func (c Count) String() string {
	return strconv.Itoa(int(c))
}

// Declare a WriteLog() function which takes any object that satisfies
// the fmt.Stringer interface as a parameter.
func WriteLog(s fmt.Stringer) {
	log.Println(s.String())
}

func main() {
	//Initialize a Book object and pass it to WriteLog()
	book := Book{"Alice in Wonderland", "Lewis Carrol"}
	WriteLog(book)

	//Initialize a Count object and pass it to WriteLog()
	count := Count(3)
	WriteLog(count)

	//  In the main function we've created different Book and Count types,
	// but passed both of them to the same WriteLog() function.
	//In turn, that calls their relevant String() functions and logs the result.
}

// Notes

//Wherever you see declaration in Go (such as a variable, function
//parameter or struct field) which has an interface type, you can
//use an object of any type so long as it satisfies the interface.

//Because this WriteLog() function uses the fmt.Stringer interface
//type in its parameter declaration, we can pass in any object that
//satisfies the fmt.Stringer interface. For example, we could pass
//either of the Book and Count types that we made earlier to the WriteLog()
//method, and the code would work OK.

//because the object being passed in satisfies the fmt.Stringer interface,
//we know that it
//has a String() string method that the WriteLog() function can safely call.
