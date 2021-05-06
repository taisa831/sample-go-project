package model

type User struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func NewUser() User {
	return User{
		ID:        1,
		Email:     "sample@example.com",
		FirstName: "Taro",
		LastName:  "Yamada",
	}
}
