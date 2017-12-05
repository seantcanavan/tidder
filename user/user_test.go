package user

import (
	"testing"
	"github.com/seantcanavan/tidder/tools"
	"fmt"
)

func TestNewUser(t *testing.T) {
	newUser, err := NewUser(tools.RandomAlphaMixed(5), tools.RandomAlphaMixed(5))
	if err != nil {
		t.Errorf("Failed to create new user: %v", err)
	}

	fmt.Println(newUser)
}