package firststruct

type User struct {
	Name     string
	password string
}

func NewUser(name, password string) User {
	return User{name, password}
}
