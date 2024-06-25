package bootcamp

import "os"

type Book struct {
	Name   string
	Author string
	Year   int
}

func Split(s string, sep rune) []string {
	t := ""
	var res []string
	for _, c := range s {
		if c == sep {
			res = append(res, t)
			t = ""
		} else {
			t += string(c)
		}
	}
	if t != "" {
		res = append(res, t)
	}
	return res
}

func stringToInt(s string) int {
	n := 0
	for _, c := range s {
		if c < '0' || c > '9' {
			return -1
		}
		n = n*10 + int(rune(c)-'0')
	}
	return n
}

func GetBooksFromCsv(path string) []*Book {
	text, _ := os.ReadFile(path)
	lines := Split(string(text), '\n')

	var name, author, year int
	t := Split(lines[0], ',')
	for i := 0; i < 3; i++ {
		if t[i] == "Name" {
			name = i
		} else if t[i] == "Author" {
			author = i
		} else {
			year = i
		}
	}

	var res []*Book
	for i := 1; i < len(lines); i++ {
		book := new(Book)
		t := Split(lines[i], ',')
		(*book).Name = t[name]
		(*book).Author = t[author]
		(*book).Year = stringToInt(t[year])
		res = append(res, book)
	}
	return res
}
