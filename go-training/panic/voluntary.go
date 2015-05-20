package main

import (
	"errors"
	"fmt"
)

// START OMIT
func main() {
	LaunchTheMissiles()
	fmt.Printf("This will not show up!")
}

func LaunchTheMissiles() {
	panic(errors.New("Out of missile error. Thanks Obama")) // HL
}

// END OMIT
