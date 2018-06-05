package main

import (
	"async"
	"fmt"
	"time"
)

var wg sync.WaitGroup

func say(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	go say("Hey")
	go say("There")
}
