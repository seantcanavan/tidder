package user

import (
	"testing"
	"github.com/seantcanavan/tidder/test"
	"reflect"
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

	cio, addErr := CreateUser(user)
	test.ErrorCheck(t, addErr, nil)
	t.Log(cio)

	gio, readErr := ReadUser(user.Id)
	test.ErrorCheck(t, readErr, nil)
	t.Log(gio)

	queried, avmErr := FromAvm(gio.Item)
	test.ErrorCheck(t, avmErr, nil)
	test.OutExpCheck(t, reflect.DeepEqual(user, queried), true)

	dio, delErr := DeleteUser(user.Id)
	test.ErrorCheck(t, delErr, nil)

	t.Log(dio)
}
