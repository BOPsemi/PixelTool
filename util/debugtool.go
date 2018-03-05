/*
Debug tool class
*/

package util

import (
	"fmt"
)

/*
DebugTool :interface of debugging tools
*/
type DebugTool interface {
	DebugPrint(obj interface{})
}

// struct definition for debugTool
type debugTool struct {
}

// NewDebugTool : initializer of debugTool
func NewDebugTool() DebugTool {
	obj := new(debugTool)

	return obj
}

/*
DebugPrint	:function for printing the object
obj interface{}
*/

func (de *debugTool) DebugPrint(obj interface{}) {
	fmt.Println(obj)
}
