package sql

import (
	"fmt"
	"testing"
)

func TestPick(t *testing.T) {
	line := "unique (`passport`)"
	res := PickUniqueFieldName(line)
	fmt.Println("结果:", res)
}

func TestIsDeleteField(t *testing.T) {
	ok := IsDeleteField("is_deleted")
	fmt.Println(ok)
}

func TestIsNameField(t *testing.T) {
	ok := IsNameField("username")
	fmt.Println(ok)
}
