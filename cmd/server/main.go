package main

import (
	"fmt"

	"github.com/bhushankp/user-service.git/internal/config"
)

func main() {
	cfg := config.LoadConfig()

	fmt.Println("sever port :", cfg.Server.Port)
}
