package main

import (
	"fmt"
	"sort"
)

type IntegerSlice []int

func (mss IntegerSlice) Len() int           { return len(mss) }                 // HL
func (mss IntegerSlice) Less(i, j int) bool { return mss[i] < mss[j] }          // HL
func (mss IntegerSlice) Swap(i, j int)      { mss[i], mss[j] = mss[j], mss[i] } // HL

func main() {
	ints := []int{6, 3, 4, 5, 1, 0, 2}
	sort.Sort(IntegerSlice(ints)) // HL
	fmt.Println(ints)
}
