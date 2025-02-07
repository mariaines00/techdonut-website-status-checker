package main

import (
	"fmt"
	"net/http"
	"runtime"
	"time"
)

func main() {

	fmt.Println("OS ", runtime.GOOS)
	fmt.Println("ARCH ", runtime.GOARCH)
	fmt.Println("CPUs ", runtime.NumCPU())
	fmt.Println("#goroutines ", runtime.NumGoroutine())

	links := []string{
		"https://reddit.com",
		"https://mariainesserra.com",
		"https://talkdesk.com",
		"https://mockymonkey.herokuapp.com/",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c { //loop using channel as range
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	resp, _ := http.Head(link)

	if resp.StatusCode >= 500 {
		c <- link
		fmt.Println(link, "might be down :(")
		return
	}

	fmt.Println(link, "is up :)")
	c <- link
}
