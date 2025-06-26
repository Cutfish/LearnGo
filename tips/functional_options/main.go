package main

import (
	"fmt"
	"time"
)

type Server struct {
	Addr           string
	Port           int
	Timeout        time.Duration
	MaxConnections int
	EnableTLS      bool
}

type Option func(s *Server)

func WithPort(port int) Option {
	return func(s *Server) {
		s.Port = port
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.Timeout = timeout
	}
}

func WithMaxConnections(max int) Option {
	return func(s *Server) {
		s.MaxConnections = max
	}
}

func WithTLS() Option {
	return func(s *Server) {
		s.EnableTLS = true
	}
}

func NewServer(addr string, opts ...Option) *Server {
	// 设置默认值
	server := &Server{
		Addr:           addr,
		Port:           8080,
		Timeout:        30 * time.Second,
		MaxConnections: 100,
		EnableTLS:      false,
	}
	for _, opt := range opts {
		opt(server)
	}

	return server
}

func main() {
	// --- 使用函数式选项模式创建 Server 实例 ---

	// 1. 使用所有默认配置
	server1 := NewServer("127.0.0.1")
	fmt.Printf("Server 1: %+v\n", server1)

	fmt.Println("---")

	// 2. 仅修改端口和超时时间
	server2 := NewServer("0.0.0.0",
		WithPort(80),
		WithTimeout(5*time.Second),
	)
	fmt.Printf("Server 2: %+v\n", server2)
	fmt.Println("---")

	// 3. 修改所有可选参数，并启用 TLS
	server3 := NewServer("localhost",
		WithPort(443),
		WithTimeout(2*time.Minute),
		WithMaxConnections(500),
		WithTLS(),
	)
	fmt.Printf("Server 3: %+v\n", server3)
}
