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

	for _, link := range links {
		checkLink(link)
	}
}

func checkLink(link string) {
	resp, _ := http.Head(link)

	if resp.StatusCode >= 500 {
		fmt.Println(link, "might be down :(")
		return
	}

	fmt.Println(link, "is up :)")
}
