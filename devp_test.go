package devp

import (
	"testing"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func Test_Stack(t *testing.T) {
	s := Stack()
	t.Logf("Stack: \n%s\n", s)
}

func Test_PrintStack(t *testing.T) {
	PrintStack("")
	PrintStack("abcdefg")
}

func Test_PrintObjectJson(t *testing.T) {
	u := User{1, "name1", 31}
	PrintObjectJson("", u)
	PrintObjectJson("u", u)
}

func Test_PrintObjectJsonWithStack(t *testing.T) {
	u := User{2, "name2", 32}
	PrintObjectJsonWithStack("", u)
	PrintObjectJsonWithStack("u", u)
}

func Test_GetObjectJson(t *testing.T) {
	u := User{3, "name3", 33}
	json := GetObjectJson("", u)
	t.Log(json)
	json = GetObjectJson("u", u)
	t.Log(json)
}

func Test_GetObjectJsonWithStack(t *testing.T) {
	u := User{4, "name4", 34}
	json := GetObjectJsonWithStack("", u)
	t.Log(json)
	json = GetObjectJsonWithStack("u", u)
	t.Log(json)
}

