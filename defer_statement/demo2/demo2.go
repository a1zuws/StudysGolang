/*You need to run main.go to print to the terminal.*/
package defer_statement

import (
	"fmt"
)
 
func Demo2(number int32) string {
	defer DeferFunc()
	
	if number%2 == 0 {
	    return "This is even number!"
	}
	
	if number%2 != 0 {
	    return "This is odd number!"
	}
	
	return "It is not clear!"
}

func TestDemo2() {
	conclusion := Demo2(9)
	fmt.Println(conclusion)
}

func DeferFunc() {
	fmt.Println("Finish!")
}
