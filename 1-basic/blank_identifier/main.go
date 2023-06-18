package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	result, _ := http.Get("https://google.com")
	body, _ := ioutil.ReadAll(result.Body)
	result.Body.Close()

	fmt.Printf("%s", body)
}
