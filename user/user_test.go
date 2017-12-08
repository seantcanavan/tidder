package user

import (
	"github.com/satori/go.uuid"
	"github.com/seantcanavan/tidder/test"
	"reflect"
	"testing"
)

func TestNewUser(t *testing.T) {
	var name = test.RandomAlphaMixed(5)
	var email = test.RandomEmail()
	user, err := New(name, email)

	test.ErrorCheck(t, err, nil)
	test.InOutExpCheck(t, email, user.Email, email)
	test.InOutExpCheck(t, name, user.Name, name)
}

func TestToAvmAndToAvmUpdateAndFromAvm(t *testing.T) {
	user := &User{
		Id:    uuid.NewV4().String(),
		Name:  "some1",
		Last:  "lastHere",
		First: "firstHere",
		Email: "at@dot.com",
	}

	updateAvm, toUpdateErr := ToAvmUpdate(user)
	test.ErrorCheck(t, toUpdateErr, nil)

	_, exists := updateAvm["Id"]
	test.OutExpCheck(t, exists, false)
	test.OutExpCheck(t, *updateAvm["Name"].Value.S, user.Name)
	test.OutExpCheck(t, *updateAvm["Last"].Value.S, user.Last)
	test.OutExpCheck(t, *updateAvm["First"].Value.S, user.First)
	test.OutExpCheck(t, *updateAvm["Email"].Value.S, user.Email)

	avm, toAvmErr := ToAvm(user)
	test.ErrorCheck(t, toAvmErr, nil)

	fromUser, fromErr := FromAvm(avm)
	test.ErrorCheck(t, fromErr, nil)

	test.OutExpCheck(t, reflect.DeepEqual(user, fromUser), true)
}
