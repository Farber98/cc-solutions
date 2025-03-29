package main

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"time"
)

type RateLimiter interface {
	Allow(ip string) bool
	GetState(ip string) map[string]any
}

type app struct {
	rl RateLimiter
}

func newApp(rl RateLimiter) *app {
	return &app{rl: rl}
}

func main() {
	//tb := NewTokenBucket(10, 5*time.Second, 5*time.Minute)
	//fw := NewFixedWindow(3, 20*time.Second)
	slw := NewSlidingWindowLog(5, 10*time.Second)
	app := newApp(slw)

	http.HandleFunc("/limited", func(w http.ResponseWriter, r *http.Request) {
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			log.Println("Error parsing RemoteAddr:", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Server Error"))
			return
		}

		if !app.rl.Allow(ip) {
			log.Println("Rate limit exceeded for /limited")
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte("Rate limit exceeded!"))
			return
		}
		log.Println("Request allowed for /limited")
		w.Write([]byte("Limited, don't over use me!"))
	})

	http.HandleFunc("/unlimited", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Unlimited! Let's Go"))
	})

	http.HandleFunc("/bucket", func(w http.ResponseWriter, r *http.Request) {
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			log.Println("Error parsing RemoteAddr:", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Server Error"))
			return
		}

		state := app.rl.GetState(ip)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(state)
	})

	http.ListenAndServe(":8080", nil)
}
