package balancer

import (
	"context"
	"log"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

type roundRobinLoadBalancer struct {
	idx     atomic.Int32
	servers []*Server
	client  *http.Client
	mu      sync.RWMutex
}

func NewRoundRobinLoadBalancer(svs []*Server, client *http.Client) *roundRobinLoadBalancer {
	lb := &roundRobinLoadBalancer{servers: svs, client: client}
	for _, sv := range svs {
		alive := lb.isAlive(sv)
		sv.mu.Lock()
		sv.alive = alive
		sv.mu.Unlock()
		if alive {
			log.Printf("Initial health check: server %s is up", sv.url.String())
		} else {
			log.Printf("Initial health check: server %s is down", sv.url.String())
		}
	}
	go lb.healthCheck()
	return lb
}

func (lb *roundRobinLoadBalancer) healthCheck() {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		for _, sv := range lb.servers {
			sv.mu.Lock()
			alive := lb.isAlive(sv)
			tmp := sv.alive
			sv.alive = alive
			sv.mu.Unlock()

			if tmp && !alive {
				log.Printf("server %s went down", sv.url.String())
			} else if !tmp && alive {
				log.Printf("server %s went up", sv.url.String())
			}
		}
	}
}

func (lb *roundRobinLoadBalancer) GetSvProxy() http.Handler {
	var sv *Server
	found := make(chan bool, 1)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	go func() {
		sv = lb.getSvWithContext(ctx)
		if sv != nil {
			found <- true
		}
	}()

	select {
	case <-found:
		return sv.proxy
	case <-ctx.Done():
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusServiceUnavailable)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"error": "All servers are currently unavailable", "status": 503}`))
			log.Printf("Error: All servers down, returning 503 for request to %s", r.RemoteAddr)
		})
	}
}

func (lb *roundRobinLoadBalancer) getSvWithContext(ctx context.Context) *Server {
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			lb.mu.RLock()
			sv := lb.servers[lb.idx.Add(1)%int32(len(lb.servers))]
			lb.mu.RUnlock()

			sv.mu.RLock()
			alive := sv.alive
			sv.mu.RUnlock()
			if alive {
				return sv
			}
		}
	}
}

func (lb *roundRobinLoadBalancer) isAlive(server *Server) bool {
	resp, err := lb.client.Get(server.url.String())
	if err != nil || resp.StatusCode != 200 {
		return false
	}
	return true
}
