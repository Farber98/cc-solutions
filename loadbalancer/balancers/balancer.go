package balancer

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
)

type LoadBalancer interface {
	healthCheck()
	GetSvProxy() http.Handler
	isAlive(*Server) bool
}

type Server struct {
	url   *url.URL
	proxy *httputil.ReverseProxy
	alive bool
	mu    sync.RWMutex
}

func NewServer(url *url.URL, proxy *httputil.ReverseProxy) *Server {
	return &Server{url: url, proxy: proxy}
}
