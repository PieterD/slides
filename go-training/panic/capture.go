package main

import (
	"fmt"
	"log"
)

// START OMIT
func main() {
	err := DoTheThing()
	fmt.Printf("%v\n", err)
}

func DoTheThing() (err error) {
	defer capture(&err)
	panic(99999)
}

func capture(err *error) {
	i := recover()
	if i != nil {
		log.Printf("Recovered: %v", i)
		*err = fmt.Errorf("DoTheThing paniced with: %v", i)
	}
}

// END OMIT
