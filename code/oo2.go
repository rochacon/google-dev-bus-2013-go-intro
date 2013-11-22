package main

import "fmt"

// START0 OMIT
type Uid int

func (i Uid) Next() Uid {
	return i + 1
}

// STOP0 OMIT

// START1 OMIT
type User struct {
	Id   Uid
	Name string
}

// STOP1 OMIT

func (u *User) String() string {
	return fmt.Sprintf("Id: %d Name: %s", u.Id, u.Name)
}

func main() {
	u := &User{Id: 1, Name: "Rodrigo Chacon"}
	fmt.Println(u)
}
