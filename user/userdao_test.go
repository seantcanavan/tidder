package user

import (
	"testing"
	"github.com/seantcanavan/tidder/test"
)

func TestUserTable(t *testing.T) {
	dto, dtoErr := DescribeTable()
	test.ErrorCheck(t, dtoErr, nil)

	test.OutExpCheck(t, TABLE_NAME, *dto.Table.TableName)
}

func TestUserCRUD(t *testing.T) {
	user, err := New(test.RandomAlphaMixed(10), test.RandomEmail())

	test.ErrorCheck(t, err, nil)

	_, addErr := AddUser(user)
	test.ErrorCheck(t, addErr, nil)

	dii, delErr := DeleteUser(user.Id)
	test.ErrorCheck(t, delErr, nil)

	t.Log(dii)
}
