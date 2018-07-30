package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"wjf.de/Flock"
)

func main() {
	println("Enter Mode:")
	in := input()
	//in := "Node"
	//in := "Manager"
	switch in {
	case "Manager":
		url, router := Flock.Manager([]byte{127, 0, 0, 1}, 5000, "NoName")
		fmt.Println(Flock.MgrInfos)
		http.ListenAndServe(url, router)
		break
	case "Node":
		authKey := input()
		//authKey := "c33d6aa2-fe23-402d-a65a-b973a708ebd3"
		manager, err := Flock.ConnectToManager([]byte{127, 0, 0, 1}, 5000, authKey)
		if err != nil {
			println("Access denied")
			return
		}
		thisNode, _ := manager.Node([]byte{127, 0, 0, 1}, 50001, "FirstNode", "SomeNode", "UDP")
		fmt.Println(thisNode)

		allNodes, _ := manager.GetConnectedNodes()
		fmt.Println(allNodes)

		thisNode.Logout(&manager)

		allNodes, _ = manager.GetConnectedNodes()
		fmt.Println(allNodes)

		break
	}
}

func input() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}
