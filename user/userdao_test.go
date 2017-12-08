package user

import (
	"github.com/satori/go.uuid"
	"github.com/seantcanavan/tidder/test"
	"reflect"
	"testing"
)

func TestUserTable(t *testing.T) {
	dto, dtoErr := DescribeTable()
	test.ErrorCheck(t, dtoErr, nil)

	test.OutExpCheck(t, TABLENAME, *dto.Table.TableName)
}

func TestUserCRUD(t *testing.T) {
	user, err := New(test.RandomAlphaMixed(10), test.RandomEmail())
	user.First = "sean"
	user.Last = "canavan"

	test.ErrorCheck(t, err, nil)

	//C
	cio, addErr := CreateUser(user)
	test.ErrorCheck(t, addErr, nil)
	t.Log("CreateUserOutput:")
	t.Log(cio)

	//U
	user.First = "newFirst"
	user.Last = "newLast"
	user.Email = "new@email.com"
	user.Id = uuid.NewV4().String()
	user.Name = "newUser"

	uuo, updateErr := UpdateUser(user)
	test.ErrorCheck(t, updateErr, nil)
	t.Log("UpdateUserOutput:")
	t.Log(uuo)

	//R
	gio, readErr := ReadUser(user.Id)
	test.ErrorCheck(t, readErr, nil)
	t.Log("ReadUserOutput:")
	t.Log(gio)

	queried, avmErr := FromAvm(gio.Item)
	test.ErrorCheck(t, avmErr, nil)
	test.OutExpCheck(t, reflect.DeepEqual(user, queried), true)

	//D
	dio, delErr := DeleteUser(user.Id)
	test.ErrorCheck(t, delErr, nil)

	t.Log("DeleteUserOutput:")
	t.Log(dio)
}
