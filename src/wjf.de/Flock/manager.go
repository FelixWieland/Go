package Flock

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/google/uuid"

	"github.com/julienschmidt/httprouter"
)

//MgrInfos Informations about the active Manager
var MgrInfos ManagerInformations

//ConnectedNodes Slice of informations about all connected Nodes
var ConnectedNodes []NodeInformations

// - /ADDNODE
func addNode(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	urlMap := parseURL(req.URL.String())

	IP := parseToBytes(urlMap["IP"])
	NodePort, _ := strconv.Atoi(urlMap["Port"])
	sNodeName := urlMap["Name"]
	sAuthKey := urlMap["AuthKey"]

	if sAuthKey != MgrInfos.AuthKey {
		w.Write([]byte("Access denied"))
		println("Tried do addNode with Authkey: " + sAuthKey)
		return
	}

	nodeInfos := NodeInformations{
		uuid.New().String(),
		sNodeName,
		IP,
		NodePort,
		[]string{""},
	}

	ConnectedNodes = append(ConnectedNodes, nodeInfos)

	infos, _ := json.Marshal(struct {
		ID    string
		Name  string
		IP    string
		Port  int
		Types []string
	}{
		nodeInfos.ID,
		nodeInfos.Name,
		byteArrToString(nodeInfos.IP),
		nodeInfos.Port,
		nodeInfos.Types,
	})

	w.Write([]byte(infos))
	print("Added Node: " + nodeInfos.ID)
}

// - /CONNECT
func incomingConnection(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	urlMap := parseURL(req.URL.String())
	sAuthKey := urlMap["AuthKey"]

	if sAuthKey != MgrInfos.AuthKey {
		w.Write([]byte("Access denied"))
		println(sAuthKey)
		return
	}
	infos, _ := json.Marshal(struct {
		ID      string
		Name    string
		IP      string
		Port    int
		AuthKey string
	}{
		MgrInfos.ID,
		MgrInfos.Name,
		byteArrToString(MgrInfos.IP),
		MgrInfos.Port,
		MgrInfos.AuthKey,
	})
	w.Write(infos)
	println("Incoming connection from: " + req.URL.String())
}

//getConnectedNodes Outputs all Connected Nodes
func getConnectedNodes(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	urlMap := parseURL(req.URL.String())
	sAuthKey := urlMap["AuthKey"]

	if sAuthKey != MgrInfos.AuthKey {
		w.Write([]byte("Access denied"))
		println("Tried do getConnectedNodes with Authkey: " + sAuthKey)
		return
	}

	type connectedNodesString struct {
		Nodes []nodeInformationsString
	}
	allNodes := connectedNodesString{}
	for index := range ConnectedNodes {
		nodeString := nodeInformationsString{
			ConnectedNodes[index].ID,
			ConnectedNodes[index].Name,
			byteArrToString(ConnectedNodes[index].IP),
			ConnectedNodes[index].Port,
			ConnectedNodes[index].Types,
		}
		allNodes.Nodes = append(allNodes.Nodes, nodeString)
	}
	allNodesJSON, _ := json.Marshal(allNodes)
	w.Write(allNodesJSON)
}

//logoutNode loggs the current Node out
func logoutNode(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	urlMap := parseURL(req.URL.String())
	sAuthKey := urlMap["AuthKey"]
	sNodeID := urlMap["NodeID"]

	if sAuthKey != MgrInfos.AuthKey {
		w.Write([]byte("Access denied"))
		println("Tried do logout a Node with Authkey: " + sAuthKey)
		return
	}

	w.Write([]byte(strconv.FormatBool(removeNode(sNodeID))))
}

func setManagerInformations(ManagerIP []byte, ManagerPort int, ManagerName string) {
	MgrInfos = ManagerInformations{
		uuid.New().String(),
		ManagerName,
		ManagerIP,
		ManagerPort,
		uuid.New().String(),
	}
}

func valueOfParam(param string) string {
	return strings.SplitAfter(param, "=")[1]
}

func removeNode(sNodeID string) bool {
	for index := range ConnectedNodes {
		if ConnectedNodes[index].ID == sNodeID {
			ConnectedNodes = append(ConnectedNodes[:index], ConnectedNodes[index+1:]...)
			return true
		}
	}
	return false
}
