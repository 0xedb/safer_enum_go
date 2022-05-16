package main

import "fmt"

// Flag is a flag
type Flag struct {
	num int
}

func (f Flag) String() string {
	return fmt.Sprintf("%d", f.num)
}

var (
	// Read flag
	Read = Flag{1 << 0}

	// Write flag
	Write = Flag{1 << 1}

	// Execute flag
	Execute = Flag{1 << 2}
)

func isReadOrWrite(f Flag) bool {
	return f.num == Read.num || f.num == Write.num
}

func main() {
	fmt.Println(isReadOrWrite(Read))
	fmt.Println(isReadOrWrite(Write))
	fmt.Println(isReadOrWrite(Execute))
}
