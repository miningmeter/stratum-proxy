package lumerinlib

import (
	"fmt"
	"testing"
)

func TestLibFunctions(t *testing.T) {

	fmt.Printf(FileLine() + "\n")
	fmt.Printf(Funcname() + "\n")

}
