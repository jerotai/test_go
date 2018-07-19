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
	
	if len(args) > 2 {
		os.Setenv("ROUTER_RSA_OPEN", args[2])
		fmt.Println(args[2])
	}
	
	api.NewApiService().Start()
}