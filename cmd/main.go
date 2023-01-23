package main

import (
	"clockifyClientApi/client"
	"fmt"
	"log"
	"time"
)

const apiKey string = "NjA1NGJiZmEtYmMzMC00OWM3LWI5NjEtZTBiYmFiMjU2ZTVk"
const workspaceId string = "63cebe1407b500028bae0e06"

func main() {
	clockifyClient, err := client.NewClient(time.Second*15, apiKey, workspaceId)
	if err != nil {
		log.Fatal(err)
	}

	clients, err := clockifyClient.GetClients()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(clients)

}
