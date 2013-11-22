package main

import "fmt"

// START0 OMIT
type Stringer interface {
	String() string
}

// STOP0 OMIT

type User struct {
	Id   Uid
	Name string
}

// START1 OMIT
func (u *User) String() string {
	return fmt.Sprintf("Id: %d Name: %s", u.Id, u.Name)
}

// STOP1 OMIT

func main() {
	// START2 OMIT
	u := &User{Id: 1, Name: "Rodrigo Chacon"}
	fmt.Println(u)
	// STOP2 OMIT
}
