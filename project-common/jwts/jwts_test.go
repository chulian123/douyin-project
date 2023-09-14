package jwts

import (
	"fmt"
	"testing"
)

func TestParseToken(t *testing.T) {
	str := CreateToken(1)
	fmt.Println(str)
	s := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjM3MjM2MywiZXhwIjoxNjk0NjE0NjgyfQ.7-pPlqC0bzrzZe5RMY3ExkOq8bqjWWmVfuO93bMxXIk"
	auth, valid := Parse(s)
	fmt.Println(auth)
	fmt.Println(valid)
}
