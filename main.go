package main

import (
	"github.com/The-EpaG/Go-BasicServer/cmd"
	"github.com/The-EpaG/Go-BasicServer/pkg/config"
)

func main() {
	config.Init()
	cmd.Execute()
}
