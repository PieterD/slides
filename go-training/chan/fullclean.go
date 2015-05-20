package main

import (
	"log"
	"math/rand"
	"sync"
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

// START OMIT
func GetThingConcurrently(url string, ch chan<- Response, kill <-chan struct{}, wg *sync.WaitGroup) { // HL
	defer wg.Done() // HL
	resp := GetTheThing(url)
	select {
	case ch <- resp:
		log.Printf("thing %s got!", url)
	case <-kill:
		log.Printf("killed %s", url)
	}
}

func GetAllTheThings(list []string) (respList []Response) {
	var wg sync.WaitGroup // HL
	wg.Add(len(list))     // HL
	defer wg.Wait()       // HL

	kill := make(chan struct{})
	defer close(kill)
	ch := make(chan Response)

	for _, url := range list {
		go GetThingConcurrently(url, ch, kill, &wg)
		// ...
		// END OMIT
	}

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
