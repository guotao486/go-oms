package test

import (
	"fmt"
	"oms/pkg/util"
	"testing"
)

// go test -v -run TestHashPassword .\test\hash_test.go
func TestHashPassword(t *testing.T) {
	password := "$2a$10$J.x6Ri1L5Byq9GO.MYxoQ.szmwganWzxH9HcTbub7KKfMXRMOFEqm"
	salt := "vaY5ef5WII"

	p := "password" + salt

	fmt.Printf("password: %v\n", password)
	b := util.CheckPasswordHash(password, p)
	fmt.Printf("b: %v\n", b)
	b = util.CheckPasswordHash(p, password)
	fmt.Printf("b: %v\n", b)

}
