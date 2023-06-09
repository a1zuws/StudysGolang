/*You need to run main.go to print to the terminal.*/
package error_handling

import (
	"fmt"
)

type borderException struct {
	parameter int
	message string
}

func (b borderException) Error() string {
	return fmt.Sprintf("%d---%s", b.parameter, b.message)
}

func Predict2(predict int) (string, error) {
	if predict < 1 || predict > 100 {
		return "" , &borderException{predict, "Out of bounds!"}
	}
	return "You know", nil
}
