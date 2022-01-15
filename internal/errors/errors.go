package errors

import (
	"fmt"
	"os"
)

func Blask(err error) {
	fmt.Println("[err]", err)
	os.Exit(1)
}
