package models

type User struct {
	IsAdmin  bool
	IsBanned bool
}

func NewUser(isAdmin bool, isBanned bool) *User {
	return &User{
		IsAdmin:  isAdmin,
		IsBanned: isBanned,
	}
}
