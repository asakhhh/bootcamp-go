package bootcamp

import (
	"bootcamp/firststruct"

	"github.com/alem-platform/ap"
)

func PrintString(s string) {
	for _, c := range s {
		ap.PutRune(c)
	}
}

func PrintUserInfo(u firststruct.User) {
	yesno := "yes"
	if u.PasswordReliability() == "undefined" {
		yesno = "no"
	}
	PrintString("Name: " + u.Name + "\nHasPassword: " + yesno + "\n")
}

// func main() {
// 	userAlem := firststruct.NewUser("Alem", "hello.alem")
// 	PrintUserInfo(userAlem)
// 	// Output:
// 	// Name: Alem
// 	// HasPassword: yes

// 	userDias := firststruct.NewUser("Dias", "")
// 	PrintUserInfo(userDias)
// 	// Output:
// 	// Name: Dias
// 	// HasPassword: no
// }
