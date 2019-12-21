package main

import (
	"fmt"
	"log"
	"os"

	"github.com/vskut/twigo/pkg/gateway"
)

func main() {
	log.Println("Start client...")

	gw := gateway.NewRestGateway()

	endpoint := fmt.Sprintf("%s:%s", os.Getenv("SERVER_HOST"), os.Getenv("SERVER_PORT"))
	gwAddr := fmt.Sprintf("%s:%s", os.Getenv("GATEWAY_HOST"), os.Getenv("GATEWAY_PORT"))
	if err := gw.Run(endpoint, gwAddr); err != nil {
		log.Panicf("failed to listen: %v", err)
	}
}
