package main

import "fmt"

// START OMIT
func main() {
	ZhuLiDoTheThing()
	fmt.Println("Returning now.")
}

func ZhuLiDoTheThing() {
	defer PrintPanic() // HL
	panic("Panic, world!")
}

func PrintPanic() {
	i := recover() // HL
	if i != nil {
		fmt.Printf("recovered: %v\n", i)
	}
}

// END OMIT
