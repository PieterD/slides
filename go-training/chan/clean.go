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
	time.Sleep(1000 * time.Millisecond)
}

func GetTheThing(url string) Response {
	time.Sleep(time.Duration(rand.Int()%1000) * time.Millisecond)
	return Response("response:" + url)
}

// START OMIT
func GetThingConcurrently(url string, ch chan<- Response, kill <-chan struct{}) { // HL
	resp := GetTheThing(url)
	select { // HL
	case ch <- resp:
		log.Printf("thing %s got!", url)
	case <-kill: // HL
		log.Printf("killed %s", url) // HL
	} // HL
}

func GetAllTheThings(list []string) (respList []Response) {
	kill := make(chan struct{}) // HL
	defer close(kill)           // HL
	ch := make(chan Response)

	for _, url := range list {
		go GetThingConcurrently(url, ch, kill)
	}
	// ...
	// END OMIT

	after := time.After(500 * time.Millisecond)
	for {
		select {
		case <-after:
			return
		case resp := <-ch:
			respList = append(respList, resp)
			if len(respList) == len(list) {
				return
			}
		}
	}
	return
}
