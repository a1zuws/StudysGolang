/*You need to run main.go to print to the terminal.*/
package error_handling

import (
	"fmt"
	"os"
)

func Demo1() {
	f,err := os.Open("demo1.txt")
	
	if err != nil{
	    fmt.Println("File not found!")
	} else {
	    fmt.Println(f.Name())
	
	}

}
