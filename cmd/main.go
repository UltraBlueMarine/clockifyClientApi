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

	//clients, err := clockifyClient.GetClients()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//for _, i := range clients {
	//	fmt.Println(i.Info())
	//}

	//client1, err := clockifyClient.GetClientById("63cef9708206496ce41c752c")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(client1.Info())

	//err = clockifyClient.InsertNewClient("Lena", "Czesc!")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//clients, err := clockifyClient.GetClients()
	//fmt.Println(clients)
	//err = clockifyClient.UpdateClient("63d067f224158e30cec49cdf", "Good Morning!", false)
	err = clockifyClient.DeleteClient("63d068a224158e30cec49e06")
	clients, err := clockifyClient.GetClients()
	fmt.Println(clients)

}
