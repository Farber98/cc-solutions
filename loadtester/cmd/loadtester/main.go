package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"net/url"
	"sync"
	"time"
)

type Config struct {
	url                *url.URL
	numRequests        int
	client             *http.Client
	concurrentRequests int
}

type Result struct {
	TTFB       time.Duration
	TTLB       time.Duration
	Error      error
	StatusCode int
}

func parseConfig() (*Config, error) {
	urlFlag := flag.String("u", "http://localhost:8081", "provide url to execute load testing against")
	nFlag := flag.Int("n", 10, "provide number of request to make")
	concurrencyFlag := flag.Int("c", 1, "concurrent requests")
	flag.Parse()
	url, err := url.Parse(*urlFlag)
	if err != nil {
		return nil, err
	}

	return &Config{
		url:                url,
		numRequests:        *nFlag,
		client:             &http.Client{},
		concurrentRequests: *concurrencyFlag,
	}, nil
}

func makeRequest(cfg *Config) Result {
	start := time.Now()
	resp, err := cfg.client.Get(cfg.url.String())
	if err != nil {
		return Result{Error: err}
	}
	ttfb := time.Since(start)
	defer resp.Body.Close()
	bytesRead, _ := io.Copy(io.Discard, resp.Body)
	_ = bytesRead
	ttlb := time.Since(start)

	return Result{
		TTFB:       ttfb,
		TTLB:       ttlb,
		StatusCode: resp.StatusCode,
	}
}

func worker(cfg *Config, jobs chan struct{}, results chan Result) {
	for range jobs {
		results <- makeRequest(cfg)
	}
}

func main() {
	cfg, err := parseConfig()
	if err != nil {
		log.Fatal(err)
	}

	jobs := make(chan struct{}, cfg.numRequests)
	results := make(chan Result, cfg.numRequests)
	numWorkers := cfg.concurrentRequests

	wg := sync.WaitGroup{}
	for range numWorkers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(cfg, jobs, results)
		}()
	}

	for range cfg.numRequests {
		jobs <- struct{}{}
	}

	go func() {
		close(jobs)
		wg.Wait()
		close(results)
	}()

	// Track metrics
	var totalTTFB, totalTTLB time.Duration
	var successes, failures int

	// Track HTTP status codes
	statusCounts := make(map[int]int)
	// For min/max calculation
	minTTLB := time.Duration(1<<63 - 1)
	var maxTTLB time.Duration

	for r := range results {
		if r.Error != nil {
			failures++
			continue
		}
		statusCounts[r.StatusCode]++
		successes++
		totalTTFB += r.TTFB
		totalTTLB += r.TTLB
		if r.TTLB < minTTLB {
			minTTLB = r.TTLB
		}
		if r.TTLB > maxTTLB {
			maxTTLB = r.TTLB
		}
	}

	log.Printf("\nSummary:\n")
	log.Printf("Total requests: %d\n", cfg.numRequests)
	log.Printf("Successful requests: %d\n", successes)
	log.Printf("Failed requests: %d\n", failures)

	log.Printf("\nStatus code counts:")
	for code, count := range statusCounts {
		log.Printf("  %d: %d\n", code, count)
	}

	if successes > 0 {
		avgTTFB := totalTTFB / time.Duration(successes)
		avgTTLB := totalTTLB / time.Duration(successes)
		log.Printf("\nTiming (among successful requests):")
		log.Printf("  Min TTLB: %v\n", minTTLB)
		log.Printf("  Max TTLB: %v\n", maxTTLB)
		log.Printf("  Avg TTFB: %v\n", avgTTFB)
		log.Printf("  Avg TTLB: %v\n", avgTTLB)

		// Simple requests-per-second calculation
		requestTime := float64(totalTTLB) / float64(time.Second)
		rps := float64(successes) / requestTime
		log.Printf("  Requests/sec: %.2f\n", rps)
	} else {
		log.Println("No successful requests, so no timing data available.")
	}
}
