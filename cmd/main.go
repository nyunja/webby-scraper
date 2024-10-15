package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	response, err := http.Get("https://jiji.co.ke/")

	if err != nil {
		fmt.Print("Error extracting response: ", err)
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading body: ", err)
	}
	fmt.Println(string(body))
}
