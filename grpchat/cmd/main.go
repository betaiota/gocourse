package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/betaiota/grpchat/pkg/client"
	"github.com/betaiota/grpchat/pkg/server"
)

func main() {
	serverFlag := flag.Bool("s", false, "Start as server")
	usernameFlag := flag.String("user", "", "User name (optional, will prompt if not provided)")
	passwordFlag := flag.String("password", "", "User password (optional, will prompt if not provided)")
	portFlag := flag.String("port", "", "Port for server/client (optional, will prompt if not provided)")
	urlFlag := flag.String("url", "", "Server address (optional, will prompt if not provided)")
	redisFlag := flag.String("redis", "localhost:6379", "Redis address (e.g., localhost:6379). Leave empty to disable Redis")

	flag.Parse()

	if *serverFlag {
		fmt.Println("There it begins...")
		server.CreateChatServer(*portFlag, *redisFlag)
	} else {
		if *urlFlag == "" || *portFlag == "" || *usernameFlag == "" {
			config, err := client.ShowMenu()
			if err != nil {
				log.Fatalf("Error in menu: %v", err)
			}
			client.CreateChatClient(config.Username, config.Password, config.ServerURL, config.ServerPort)
		} else {
			password := *passwordFlag
			if password == "" {
				password = "admin"
			}
			client.CreateChatClient(*usernameFlag, password, *urlFlag, *portFlag)
		}
	}
}
