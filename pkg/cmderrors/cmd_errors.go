package cmderrors

import (
	"fmt"
	"os"
)

func ExitBadly(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
