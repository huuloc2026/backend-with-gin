package main

import (
	"github.com/huuloc2026/go-backend/internal/routers"
)

func main() {
	r := routers.NewRoute()
	r.Run()
}
