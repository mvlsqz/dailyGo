package foo

type user struct {
	First, Last, password string
}

// while this is valid, but not encouraged, lets find how to return an interface
func NewUser(name, last, password string) user {
	usr := user{First: name, Last: last, password: password}
	return usr
}
