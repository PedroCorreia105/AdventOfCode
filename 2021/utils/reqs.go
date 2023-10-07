package utils

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func MakeRequest(token string, year int, day int) string {
	url := fmt.Sprintf("https://adventofcode.com/%v/day/%v/input", year, day)
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		panic(err)
	}
	
	req.Header.Set("Cookie", "session=" + token)
	
	response, err := retryableHTTPRequest(req, 5)

	if err != nil {
		panic(err)
	}

	body, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	return string(body)
}

func retryableHTTPRequest(request *http.Request, retries int) (*http.Response, error) {
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	for i := 0; i < retries; i++ {
		resp, err := client.Do(request)
		if err != nil {
			time.Sleep(time.Second * 2) // wait before retrying
			continue
		}

		return resp, nil
	}

	return nil, fmt.Errorf("maximum retries exceeded")
}