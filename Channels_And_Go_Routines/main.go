package main

import (
	"fmt"
	"net/http"
	"time"
)

// main routine
func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	/*the Problem here is that we need to wait for each response from server
	for _, link := range links {
		checkLink(link)
	}*/

	// We can fix it with brand new routine on each loop with channel:

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c) // go - creates a new thread
		// fmt.Println(<-c)
	}

	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c)
	// }

	// for l := range c { // Wait for the channel to return some value
	// 	time.Sleep(5 * time.Second) // 5 seconds main routine sleeps
	// 	go checkLink(l, c)
	// }

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			fmt.Println("************************")
			checkLink(link, c)
		}(l) // () -  we are imediatly call this func
	}

	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)

	// fmt.Println(<-c) - program frizzes
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Printf("%s might be down!\n", link)
		c <- link
		return
	}

	fmt.Printf("%s is up!\n", link)
	c <- link
}

/* Concurrency ( one go routine per CPU with Go Scheduler) - is not Parallelism  */
