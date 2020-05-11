package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

func main() {

	client := &http.Client{}

	//Concurrency
	//Use Wait Group to wait
	var wg sync.WaitGroup

	//Create a channel to pass data
	// c := make(chan string)

	wg.Add(2)
	go sendHTTPReq(client, &wg)
	go sendHTTPReq(client, &wg)
	// go testProcessWithChannel(3*time.Second, c)
	// go testProcessWithChannel(1*time.Second, c)

	// Wait for routine to complete
	//Wait group
	wg.Wait()

	//Channel
	// for i := 0; i < 2; i++ {
	// 	fmt.Println(<-c)
	// }

	fmt.Println("done")

}

func sendHTTPReq(client *http.Client, wg *sync.WaitGroup) {
	fmt.Println("sending...")

	//Form url
	method := "POST"
	rawurl := "http://localhost:8080"
	resource := "/user/"
	data := url.Values{}
	data.Set("name", "foo")
	data.Set("surname", "bar")
	u, _ := url.ParseRequestURI(rawurl)
	u.Path = resource
	urlStr := u.String()
	r, _ := http.NewRequest(method, urlStr, strings.NewReader(data.Encode()))

	resp, err := client.Do(r)
	if err != nil {
		log.Fatal("ERROR:", err)
	}
	defer resp.Body.Close()

	fmt.Println(resp.StatusCode)
	wg.Done()

}

func testProcessWithChannel(d time.Duration, c chan string) {
	fmt.Println("starting ...")
	msg := "Done sleep" + string(d)
	fmt.Println(msg)
	time.Sleep(d)

	c <- "msg"

}
