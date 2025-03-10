package db

import "fmt"

type User struct {
	ID int
}

func LoadUser(id int) (User, error) {
	if id < 0 {
		return User{}, fmt.Errorf("invalid negative user id '%d'", id)
	}
	if id > 100 {
		return User{}, fmt.Errorf("invalid user id '%d'", id)
	}

	return User{ID: id}, nil
}

func (u User) String() string {
	return fmt.Sprintf("this user is id: %d", u.ID)
}
