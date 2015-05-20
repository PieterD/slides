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

// START OMIT

func GetThingConcurrently(url string, ch chan<- Response) { // HL
	resp := GetTheThing(url)         // HL
	log.Printf("thing %s got!", url) // HL
	ch <- resp                       // HL
}

func GetAllTheThings(list []string) (respList []Response) {
	ch := make(chan Response) // HL
	for _, url := range list {
		go GetThingConcurrently(url, ch) // HL
	}
	for resp := range ch { // HL
		respList = append(respList, resp)
		if len(respList) == len(list) {
			return
		}
	}
	return
}

// END OMIT
