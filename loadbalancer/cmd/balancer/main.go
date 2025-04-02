package main

import (
	"flag"
	balancer "loadbalancer/balancers"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type config struct {
	serverList          []string
	clientTimeout       time.Duration
	healthCheckInterval time.Duration
}

func parseConfig() *config {
	servers := flag.String("servers", "http://localhost:8081,http://localhost:8082", "Comma-separated list of server URLs")
	clientTimeout := flag.Int("clientTimeout", 5, "HTTP client timeout in seconds")
	healthCheckInterval := flag.Int("healthCheckInterval", 5, "Health check interval in seconds")
	flag.Parse()

	serverList := strings.Split(*servers, ",")
	if len(serverList) == 0 {
		log.Fatal("No servers provided.")
	}

	return &config{
		serverList:          serverList,
		clientTimeout:       time.Duration(*clientTimeout) * time.Second,
		healthCheckInterval: time.Duration(*healthCheckInterval) * time.Second}
}

func main() {
	cfg := parseConfig()
	var svs []*balancer.Server
	for _, s := range cfg.serverList {
		targetURL, err := url.Parse(s)
		if err != nil {
			log.Fatalf("failed to parse target URL: %v", err)
		}
		sv := balancer.NewServer(targetURL)
		svs = append(svs, sv)
	}

	lb := balancer.NewRoundRobinLoadBalancer(svs, cfg.clientTimeout, cfg.healthCheckInterval)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		lb.GetSvProxy().ServeHTTP(w, r)
	})

	log.Println("Load balancer listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
