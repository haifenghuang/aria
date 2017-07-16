package interpreter

import (
	"fmt"
)

// IO.puts(any)
// Write a string to stdout and a newline.
func io_puts(args ...DataType) (DataType, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("IO.puts expects exactly 1 argument")
	}

	fmt.Println(args[0].Inspect())

	// Return a dummy string just to surpress errors,
	// as there's nothing to return.
	return &StringType{Value: ""}, nil
}

// IO.write(any)
// Write a string to stdout.
func io_write(args ...DataType) (DataType, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("IO.puts expects exactly 1 argument")
	}

	fmt.Print(args[0].Inspect())

	return &StringType{Value: ""}, nil
}