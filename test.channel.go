package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	startWatch := time.Now()
	channel1 := make(chan string)
	servers := []string{
		"http://platzi.com",
		"http://google.com",
		"http://facebook.com",
		"http://instagram.com",
	}

	for _, server := range servers {
		go checkStatusServer(server, channel1)

	}

	for index, _ := range servers {
		fmt.Println(<-channel1, index+1)
	}

	timeExec := time.Since(startWatch)

	fmt.Printf("Execution time: %s", timeExec)
}

func checkStatusServer(server string, channel chan string) {
	_, err := http.Get(server)
	if err != nil {
		//fmt.Println(server, "Is not Available =(")
		channel <- server + " Is not Available =("
	} else {
		//fmt.Println(server, "It's working OK")
		channel <- server + " It's working OK"
	}
}
