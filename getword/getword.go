package getword

import (
	"fmt"
	"io"
	"net/http"
)

func GetWord(uri string) []byte {
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		fmt.Printf("Error: Cannot form request: %v", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Error: Could not do request: %v", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error: Failed to read request body: %v", err)
	}

	return body
}
