package main

import (
	"os"
	"Stingray/api"
	"fmt"
)

func main() {
	args := os.Args
	fmt.Printf("Start Api Server!!")
	
	if len(args) > 1 {
		os.Setenv("ROUTER_PATH", args[1])
	}
	
	api.NewApiService().Start()
}