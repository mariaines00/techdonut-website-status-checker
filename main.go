package main

import (
	"fmt"
	"net/http"
	"runtime"
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

	fmt.Println(<-c) //show why this is needed to 'pause' main routine
}

func checkLink(link string, c chan string) {
	resp, _ := http.Head(link)

	if resp.StatusCode >= 500 {
		c <- "From channel: might be down"
		fmt.Println(link, "might be down :(")
		return
	}

	fmt.Println(link, "is up :)")
	c <- "From channel:yes it's up"
}
