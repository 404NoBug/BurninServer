package main

import "BurninProject/network"

func main() {
	client := network.NewClient(":8888")
	client.Run()
	select {}

}
