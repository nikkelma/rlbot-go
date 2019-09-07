// +build windows

package native

import (
	"fmt"
)

func newBridge() Bridge {
	a := getDllName()
	fmt.Printf(a)
	return nil
}
