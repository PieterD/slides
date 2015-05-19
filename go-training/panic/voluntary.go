package main

import "errors"

func main() {
	LaunchTheMissiles()
}

func LaunchTheMissiles() {
	panic(errors.New("Out of missile error. Thanks Obama"))
}
