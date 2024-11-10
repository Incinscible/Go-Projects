package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"
)

type Website struct {
	Link   string `json:"link"`
	Status bool   `json:"status"`
}

type Result struct {
	Websites []Website `json:"websites"`
}

func NewWebsite(link string) *Website {
	return &Website{
		Link:   link,
		Status: false,
	}
}

func loadWebsitesFile(filename string) ([]*Website, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var websites []*Website
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		website := NewWebsite("https://" + string(scanner.Text()))
		websites = append(websites, website)
	}

	return websites, err
}

func retryCheckWebsite(url string, retries int) bool {
	for i := 0; i < retries; i++ {
		if checkWebsite(url) {
			return true
		}
	}
	return false
}

func checkWebsite(url string) bool {
	client := http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return false
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Connection", "keep-alive")

	resp, err := client.Do(req)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	return resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusPermanentRedirect
}

func saveResultsToFile(results Result) error {
	data, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		return err
	}

	file, err := os.Create("results.json")
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	return err
}

func main() {
	fmt.Println("-- Services Checker --")
	websites, err := loadWebsitesFile("websites.txt")
	if err != nil {
		fmt.Println("error reading file")
		return
	}
	fmt.Println("[+] Reading \"websites.txt\" complete")

	var results Result

	var wg sync.WaitGroup
	semaphore := make(chan struct{}, 5)
	mu := sync.Mutex{}

	fmt.Println("[+] Checking all services...")

	for _, website := range websites {
		wg.Add(1)
		semaphore <- struct{}{}

		go func(site *Website) {
			defer wg.Done()
			defer func() { <-semaphore }()

			status := checkWebsite(site.Link)
			if !status {
				status = retryCheckWebsite(site.Link, 2)
			}
			mu.Lock()
			site.Status = status
			results.Websites = append(results.Websites, *site)
			mu.Unlock()

			// fmt.Println("URL: ", site.Link, " / Status: ", site.Status)
		}(website)
	}

	wg.Wait()

	err = saveResultsToFile(results)
	if err != nil {
		fmt.Println("[-] Error saving results")
		return
	}
	fmt.Println("[+] Results saved in \"results.txt\"")
}
