package cmd

import (
	"fmt"
	"os"
)

// ExitWithError handle exit with erro
func ExitWithError(code int, err error) {
	fmt.Fprintln(os.Stderr, "Terminal:", err)
	os.Exit(code)
}
