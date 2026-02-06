package util
import (
	"fmt"
)

func ErrorHandling(err error) {
	if err != nil {
		fmt.Println("Sorry! Unable to Connect with Database", err)
	}

}