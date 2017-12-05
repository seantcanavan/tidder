package user

import (
	"testing"
	"github.com/seantcanavan/tidder/tools"
	"fmt"
)

func TestAddUser(t *testing.T) {
	newUser, err := NewUser(tools.RandomAlphaMixed(10), tools.RandomEmail())

	fmt.Println(newUser)
	fmt.Println(err)
}