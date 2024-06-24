package firststruct

func (u User) Compare(v User) bool {
	return u.Name == v.Name && u.password == v.password
}
