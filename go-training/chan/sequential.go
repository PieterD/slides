package main

import (
	"log"
	"math/rand"
	"time"
)

// START OMIT
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

func GetAllTheThings(list []string) (respList []Response) {
	for _, url := range list {
		resp := GetTheThing(url)
		log.Printf("thing %s got!", url)
		respList = append(respList, resp)
	}
	return
}

// END OMIT
