package user

import (
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
	user := TestUser()

	//C
	cio, addErr := CreateUser(user)
	test.ErrorCheck(t, addErr, nil)
	t.Log("CreateUserOutput:")
	t.Log(cio)

	//U
	user.First = "newFirst"
	user.Last = "newLast"
	user.Email = "new@email.com"
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

func TestEmailIndex(t *testing.T) {
	user := TestUser()

	_, createErr := CreateUser(user)
	test.ErrorCheck(t, createErr, nil)

	results, queryErr := ReadUserByEmail(user.Email)
	test.ErrorCheck(t, queryErr, nil)

	test.OutExpCheck(t, len(results), 1)
	test.OutExpCheck(t, reflect.DeepEqual(results[0], user), true)

	_, deleteErr := DeleteUser(user.Id)
	test.ErrorCheck(t, deleteErr, nil)
}