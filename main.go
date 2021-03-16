package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	size := 100
	wg.Add(size)
	for i:=0;i<size;i++ {
		go sendMessage(i,wg)
	}
	wg.Wait()
}

func sendMessage(id int,wg sync.WaitGroup) {
	defer wg.Done()
	url := "http://localhost:8080/v1/message"
	method := "POST"

	payload := strings.NewReader(`{
    "Msg_id":`+strconv.Itoa(id)+`,
    "Sender":"Tom",
    "Msg":"Hello"
}`)

	client := &http.Client {}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}