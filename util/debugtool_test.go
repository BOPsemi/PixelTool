package util

import "testing"
import "github.com/stretchr/testify/assert"

// test new
func Test_NewDebugTool(t *testing.T) {
	obj := NewDebugTool()

	assert.NotNil(t, obj)
}

type TestStruct struct {
	name string
	id   int
}

func Test_DebugPrint(t *testing.T) {
	obj := NewDebugTool()

	// print string
	str1 := "hoge"
	obj.DebugPrint(str1)

	// test object
	str2 := TestStruct{
		name: "test structure",
		id:   1,
	}
	obj.DebugPrint(str2)

	// test object slices
	str3 := TestStruct{
		name: "the 2nd structure",
		id:   2,
	}

	strSlice := []TestStruct{str2, str3}
	obj.DebugPrint(strSlice)
}
