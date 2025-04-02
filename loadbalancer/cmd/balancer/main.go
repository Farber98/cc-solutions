package main

import (
	"flag"
	balancer "loadbalancer/balancers"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"
)

type app struct {
	lb balancer.LoadBalancer
}

func main() {
	servers := flag.String("servers", "http://localhost:8081,http://localhost:8082", "Comma-separated list of server URLs")
	flag.Parse()

	serverList := strings.Split(*servers, ",")
	if len(serverList) == 0 {
		log.Fatal("No servers provided.")
	}

	var svs []*balancer.Server
	client := &http.Client{Timeout: 5 * time.Second}
	for _, s := range serverList {
		targetURL, err := url.Parse(s)
		if err != nil {
			log.Fatalf("failed to parse target URL: %v", err)
		}
		sv := balancer.NewServer(targetURL, httputil.NewSingleHostReverseProxy(targetURL))
		svs = append(svs, sv)
	}

	lb := balancer.NewRoundRobinLoadBalancer(svs, client)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		lb.GetSvProxy().ServeHTTP(w, r)
	})

	log.Println("Load balancer listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
