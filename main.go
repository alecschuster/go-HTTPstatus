package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}
	//treat this just like any other value in go
	stat := make(chan string)
	//main routine does not care if child routines have work to do
	for _, link := range links {
		//when blocking call reached, second go routine spun up
		//child goroutine
		go checkLink(link, stat)
	}
	//wait for channel to receive some value infinitely
	//assign to l and immediately execute body
	for l := range stat {
		//function literal (anonymous function)
		//tell it to expect to receive l
		go func(l string) {
			time.Sleep(5 * time.Second)
			//blocking call, main routine is put to sleep
			//when message is sent, main routine wakes back up and exits
			checkLink(l, stat)
			//pass in the function
		}(l)
	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link) //blocking call, gives control back to main program
	// so other goroutines can run
	if err != nil {
		fmt.Println(link + " might be down!")
		c <- link
		return
	}
	fmt.Println(link + " is up!")
	c <- link
	return
}
