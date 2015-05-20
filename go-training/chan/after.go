package main

import (
	"log"
	"math/rand"
	"time"
)

type Response string

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	GetAllTheThings([]string{"one", "two", "three", "four", "five"})
}

func GetTheThing(url string) Response {
	time.Sleep(time.Duration(rand.Int()%1000) * time.Millisecond)
	return Response("response:" + url)
}

func GetThingConcurrently(url string, ch chan<- Response) {
	resp := GetTheThing(url)
	log.Printf("thing %s got!", url)
	ch <- resp
}

// START OMIT

func GetAllTheThings(list []string) (respList []Response) {
	ch := make(chan Response)
	for _, url := range list {
		go GetThingConcurrently(url, ch)
	}

	after := time.After(500 * time.Millisecond) // HL
	for {
		select { // HL
		case <-after: // HL
			return // HL
		case resp := <-ch:
			respList = append(respList, resp)
			if len(respList) == len(list) {
				return
			}
		} // HL
	}
	return
}

// END OMIT
