package model

type User struct {
	FirstName string
	LastName  string
	age       int
	Address
}

type Address struct {
	City string
	code int
	name string
}
