package main

import (
	"fmt"
	"io"
	"net/http"
)

const baseUrl = "http://localhost:8080/random"

func main() {
	resp, err := http.Get(baseUrl)
	if err != nil {
		fmt.Printf("response error: %v", err)
		return
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("read body error: %v", err)
		return
	}

	fmt.Println(string(data))

}
