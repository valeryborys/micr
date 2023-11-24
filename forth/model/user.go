package model

import (
	"fmt"
	"time"
)

type User struct {
	Id           Id
	Name         string
	CreationTime time.Time
}

func (u *User) Introduce() string {
	return fmt.Sprintf("My name is %v", u.Name)
}
