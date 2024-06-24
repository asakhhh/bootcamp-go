package bootcamp

import (
	"fmt"

	"bootcamp/firststruct"
)

func PrintUserInfo(u firststruct.User) {
	yesno := "yes"
	if u.PasswordReliability() == "undefined" {
		yesno = "no"
	}
	fmt.Printf("Name: %s\nHasPassword: %s\n", u.Name, yesno)
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
