package main

import (
	"fmt"
	"net/http"
)

type Result struct {
	Err      error          `json:"err"`
	Response *http.Response `json:"response"`
}

func checkStatus(done <-chan interface{}, urls ...string) <-chan Result {
	results := make(chan Result)

	go func() {
		defer close(results)
		for _, url := range urls {
			var result Result
			resp, err := http.Get(url)
			result = Result{Err: err, Response: resp}
			select {
			case <-done:
				fmt.Println("done called")
				return
			case results <- result:
				fmt.Println("result add in queue")
			}
		}
	}()
	return results
}

func main() {
	done := make(chan interface{})
	defer close(done)
	urls := []string{"https://www.google.com", "https://badhost"}
	for result := range checkStatus(done, urls...) {
		if result.Err != nil {
			fmt.Printf("error: %v\n", result.Err)
			continue
		}
	}

}
