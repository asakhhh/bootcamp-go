package bootcamp

import "github.com/alem-platform/ap"

// type Book struct {
// 	Name   string
// 	Author string
// 	Year   int
// }

func toString(n int) string {
	var res string
	minus := false
	if n < 0 {
		minus = true
		n = -n
	}

	for n > 0 {
		res = string(rune('0'+n%10)) + res
		n /= 10
	}
	if minus {
		res = "-" + res
	}

	return res
}

func trailToLen(s string, n int) string {
	for len(s) < n {
		s += " "
	}
	return s
}

func printString(s string) {
	for _, c := range s {
		ap.PutRune(c)
	}
}

func PrintBooks(books []*Book) {
	lenname, lenauthor, lenyear := 0, 0, 0
	for _, book := range books {
		if len(book.Name) > lenname {
			lenname = len(book.Name)
		}
		if len(book.Author) > lenauthor {
			lenauthor = len(book.Author)
		}
		if len(toString(book.Year)) > lenyear {
			lenyear = len(toString(book.Year))
		}
	}
	printString(trailToLen("Name", lenname) + " " + trailToLen("Author", lenauthor) + " " + "Year\n")
	for _, book := range books {
		printString(trailToLen(book.Name, lenname) + " " + trailToLen(book.Author, lenauthor) + " " + trailToLen(toString(book.Year), lenyear) + "\n")
	}
}

// func main() {
// 	books := []*Book{
// 		{Name: "The Kaizen Way", Author: "Robert Maurer", Year: 2009},
// 		{Name: "Dialogs", Author: "Plato", Year: -400},
// 		{Name: "Unknown", Author: "Unknown", Year: 10},
// 	}

// 	PrintBooks(books)
// }
