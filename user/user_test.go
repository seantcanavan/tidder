package user

import (
	"testing"
	"github.com/seantcanavan/tidder/test"
)

func TestNewUser(t *testing.T) {
	var username = test.RandomAlphaMixed(5)
	var emailAddress = test.RandomEmail()
	user, err := New(username, emailAddress)

	test.ErrorCheck(t, err, nil)
	test.InOutExpCheck(t, emailAddress, user.Email, emailAddress)
	test.InOutExpCheck(t, username, user.Name, username)
}
