package balancer

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
	"time"
)

type LoadBalancer interface {
	healthCheck(time.Duration)
	GetSvProxy() http.Handler
	isAlive(*Server) bool
}

type Server struct {
	url   *url.URL
	proxy *httputil.ReverseProxy
	alive bool
	mu    sync.RWMutex
}

func NewServer(url *url.URL) *Server {
	return &Server{url: url, proxy: httputil.NewSingleHostReverseProxy(url)}
}
