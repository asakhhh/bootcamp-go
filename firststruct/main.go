package firststruct

type User struct {
	name     string
	password string
}

func NewUser(name, password string) User {
	return User{name, password}
}
