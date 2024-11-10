package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
)

func doWork() int {
	return rand.Int()
}

func reportBack(num int) {
	snum := strconv.Itoa(num)

	host := "35.236.102.9:8080"

	resp, err := http.Get("http://" + host + "/report?payload=" + snum)
	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	fmt.Println(string(body))
}

func main() {
	reportBack(doWork())
}
