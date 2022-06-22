package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Res struct {
	Msg string `json:"message"`
}

func main() {

	response, err := http.Get("http://localhost:8000/ping")
	if err != nil {
		_ = fmt.Errorf("Error. error: %s", err.Error())
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		_ = fmt.Errorf("Failed to fetch the response body. Got error : %v", err)
	}

	var k Res

	err = json.Unmarshal(body, &k)
	if err != nil {
		_ = fmt.Errorf("Failed to unmarshal the response body. Got error : %v", err)
	}

	_, _ = fmt.Print(k)
}
