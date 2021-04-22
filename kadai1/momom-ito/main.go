package main

import (
	"fmt"
	"os"

	"github.com/momom-ito/gopherdojo-studyroom/kadai1/momom-ito/go_practice"
)

func main() {
	err := go_practice.Convpic()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}