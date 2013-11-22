package main

import "fmt"

type User struct {
	Id   int
	Name string
}

// START0 OMIT
func (u *User) String() string {
	return fmt.Sprintf("Id: %d Name: %s", u.Id, u.Name)
}

// STOP0 OMIT
func main() {
	u := &User{Id: 1, Name: "Rodrigo Chacon"}
	fmt.Println(u.String())
}
