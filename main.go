package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Todo struct represents a todo item in the JSON response
type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	// URL of the API you want to query
	// in this case we are only getting a single todo item
	apiUrl := "https://jsonplaceholder.typicode.com/todos/1"

	// Make a GET request to the API
	response, err := http.Get(apiUrl)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer response.Body.Close()

	// Read the response body into a byte slice
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// create an empty instance of the Todo struct
	var data Todo

	// can check the response code here
	// fmt.Println(string(responseBody))

	// Unmarshal the JSON response into a pointer to the empty Todo struct
	err = json.Unmarshal(responseBody, &data)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// at this point, data contains the parsed JSON
	// at can be accessed using the dot notation
	fmt.Println(data.Title)
}
