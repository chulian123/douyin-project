package encrypts

import (
	"fmt"
	"testing"
)

func TestGetUserID(t *testing.T) {
	id := GetUserID()
	fmt.Println(id)
}
