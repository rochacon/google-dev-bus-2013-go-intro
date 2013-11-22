package main

import "fmt"

type Uid int

// START0 OMIT
type Permission struct {
	perm string
}

func (up *Permission) HasPerm(perm string) bool {
	return up.perm == perm
}

type User struct {
	Id         Uid
	Name       string
	Permission // embed Permission type in User
}

func main() {
	u := &User{Id: 1, Name: "Rodrigo Chacon"}
	fmt.Println(u.HasPerm("test"))
}

// STOP0 OMIT
