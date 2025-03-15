package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
)

func fetchData(url string, wg *sync.WaitGroup) {
	defer wg.Done()

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching: ", url, "-", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", url, "-", err)
		return
	}
	fmt.Printf("Data from %s:\n%s...\n\n", url, string(body[:100]))
}

func main() {
	var urls []string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter  URLs(one per line, type 'done' to finish): ")
	for {
		scanner.Scan()
		url := scanner.Text()
		if url == "done" {
			break
		}
		urls = append(urls, url)
	}
	if len(urls) == 0 {
		fmt.Println("No URLs provided. Exiting...")
		return
	}

	var wg sync.WaitGroup

	fmt.Println("Starting parallel data fetching...\n")

	for _, url := range urls {
		wg.Add(1) //
		go fetchData(url, &wg)
	}

	wg.Wait()
	fmt.Println("All requests completed.")

}
