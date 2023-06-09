/*You need to run main.go to print to the terminal.*/
package error_handling

import (
	"errors"
	"fmt"
)

func predict(predict int) (string, error) {
	
	numberInMind := 80
	
	if predict < 1 || predict > 100 {
	    return "", errors.New("Enter a number between 1 and 100:")
        }
	
	if predict > numberInMind {
	    return "Enter a smaller number :", nil
        }
	
	if predict < numberInMind {
	    return "Enter a larger number :", nil
	}
	
	return "You know!", nil
}

func Demo3() {
	message, error := predict(80)
	fmt.Println(message)
	fmt.Println(error)
}
