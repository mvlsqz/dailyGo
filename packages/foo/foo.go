package foo

type User struct {
	First, Last, password string
}

// NewUser returns a user with the initialized values provided
func NewUser(first, last, password string) User {
	return User{
		First:    first,
		Last:     last,
		password: password,
	}
}
