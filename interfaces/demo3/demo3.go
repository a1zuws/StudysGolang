/*You need to run main.go to print to the terminal.*/
package interfaces

import (
	"fmt"
)

func verify(i interface{}) {
    number, ok := i.(int)
