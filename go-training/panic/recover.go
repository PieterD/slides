package main

import (
	"fmt"
	"log"
)

// START OMIT
func main() {
	ZhuLiDoTheThing()
	fmt.Println("Returning now.")
}

func ZhuLiDoTheThing() {
	defer logPanic() // HL
	panic("Panic, world!")
}

func logPanic() {
	i := recover() // HL
	if i != nil {
		log.Printf("recovered: %v\n", i)
	}
}

// END OMIT
