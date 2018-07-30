package Flock

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"
)

//ManagerInformations contains the Informations about a created Manager
type ManagerInformations struct {
	ID      string
	Name    string
	IP      []byte
	Port    int
	AuthKey string
}

//NodeInformations contains the Informations about a created Node
type NodeInformations struct {
	ID    string
	Name  string
	IP    []byte
	Port  int
	Types []string
}

type nodeInformationsString struct {
	ID    string
	Name  string
	IP    string
	Port  int
	Types []string
}

type managerInformationsString struct {
	ID      string
	Name    string
	IP      string
	Port    int
	AuthKey string
}

//Node Creates a new Node
func Node(sManagerIP string, sManagerPort string, sManagerID string) {

}

//Node Creates a New Node
func (Manager *ManagerInformations) Node(NodeIP []byte, NodePort int, sNodeName string, sNodeType string, sNodeConnType string) (NodeInformations, error) {

	url := "http://" + byteArrToString(Manager.IP) + ":" + strconv.Itoa(Manager.Port) + "/ADDNODE/" + "&IP=" + byteArrToString(NodeIP) + "&Port=" + strconv.Itoa(NodePort) + "&Name=" + sNodeName + "&AuthKey=" + Manager.AuthKey
	rs, err := http.Get(url)
	// Process response
	if err != nil {
		panic(err) // More idiomatic way would be to print the error and die unless it's a serious error
	}
	defer rs.Body.Close()

	bodyBytes, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		panic(err)
	}

	bodyString := string(bodyBytes)
	if bodyString == "Access denied" {
		return NodeInformations{}, errors.New("Access denied, wrong AuthKey")
	}

	nodeInfosString := nodeInformationsString{}
	json.Unmarshal(bodyBytes, &nodeInfosString)
	nodeInfos := NodeInformations{
		nodeInfosString.ID,
		nodeInfosString.Name,
		parseToBytes(nodeInfosString.IP),
		nodeInfosString.Port,
		nodeInfosString.Types,
	}
	return nodeInfos, nil
}

//Manager Creates a new Manager
func Manager(ManagerIP []byte, ManagerPort int, ManagerName string) (ip string, router *httprouter.Router) {

	setManagerInformations(ManagerIP, ManagerPort, ManagerName)

	router = httprouter.New()
	router.GET("/ADDNODE/*filepath", addNode)
	router.GET("/CONNECT/*filepath", incomingConnection)
	router.GET("/GETCONNECTEDNODES/*filepath", getConnectedNodes)
	router.GET("/LOGOUT/*filepath", logoutNode)
	return byteArrToString(ManagerIP) + ":" + strconv.Itoa(ManagerPort), router
}

//ConnectToManager connects to a Manager
func ConnectToManager(ManagerIP []byte, ManagerPort int, AuthKey string) (ManagerInformations, error) {
	url := "http://" + byteArrToString(ManagerIP) + ":" + strconv.Itoa(ManagerPort) + "/CONNECT/&AuthKey=" + AuthKey
	rs, err := http.Get(url)
	// Process response
	if err != nil {
		panic(err)
	}
	defer rs.Body.Close()

	bodyBytes, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		panic(err)
	}

	bodyString := string(bodyBytes)
	if bodyString == "Access denied" {
		return ManagerInformations{}, errors.New(bodyString)
	}

	mgrInfosStrIP := managerInformationsString{}
	json.Unmarshal(bodyBytes, &mgrInfosStrIP)

	byteIP := parseToBytes(mgrInfosStrIP.IP)
	mgrInfos := ManagerInformations{
		mgrInfosStrIP.ID,
		mgrInfosStrIP.Name,
		byteIP,
		mgrInfosStrIP.Port,
		mgrInfosStrIP.AuthKey,
	}
	return mgrInfos, nil
}

//GetConnectedNodes Returns a Slice with all connected Nodes
func (Manager *ManagerInformations) GetConnectedNodes() ([]NodeInformations, error) {
	type connectedNodesString struct {
		Nodes []nodeInformationsString
	}

	url := "http://" + byteArrToString(Manager.IP) + ":" + strconv.Itoa(Manager.Port) + "/GETCONNECTEDNODES/&AuthKey=" + Manager.AuthKey
	rs, err := http.Get(url)

	// Process response
	if err != nil {
		panic(err)
	}
	defer rs.Body.Close()

	bodyBytes, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		panic(err)
	}

	bodyString := string(bodyBytes)
	if bodyString == "Access denied" {
		return []NodeInformations{}, errors.New(bodyString)
	}

	allNodes := []NodeInformations{}
	allNodesString := connectedNodesString{}
	json.Unmarshal(bodyBytes, &allNodesString)

	for index := range allNodesString.Nodes {
		node := NodeInformations{
			allNodesString.Nodes[index].ID,
			allNodesString.Nodes[index].Name,
			parseToBytes(allNodesString.Nodes[index].IP),
			allNodesString.Nodes[index].Port,
			allNodesString.Nodes[index].Types,
		}
		allNodes = append(allNodes, node)
	}
	return allNodes, nil
}

func (Node *NodeInformations) Logout(Manager *ManagerInformations) (bool, error) {
	url := "http://" + byteArrToString(Manager.IP) + ":" + strconv.Itoa(Manager.Port) + "/LOGOUT/&AuthKey=" + Manager.AuthKey + "&NodeID=" + Node.ID
	rs, err := http.Get(url)

	// Process response
	if err != nil {
		panic(err)
	}
	defer rs.Body.Close()

	bodyBytes, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		panic(err)
	}

	bodyString := string(bodyBytes)
	if bodyString == "Access denied" {
		return false, errors.New(bodyString)
	} else if bodyString == "true" {
		return true, nil
	}
	return false, errors.New("Logout Node failed")
}

func byteArrToString(barr []byte) string {
	str := ""
	for index := range barr {
		points := ""
		if index < len(barr)-1 {
			points = "."
		}
		str = str + strconv.Itoa(int(barr[index])) + points
	}
	return str
}

func parseToBytes(sIP string) []byte {
	ip := []byte{}
	sIPArr := strings.SplitAfter(sIP, ".")
	for x := range sIPArr {
		intVal, _ := strconv.Atoi(strings.Replace(sIPArr[x], ".", "", -1))
		ip = append(ip, byte(intVal))
	}
	return ip
}

func parseURL(url string) map[string]string {
	urlMap := make(map[string]string)
	splitted1 := strings.SplitAfter(url, "&")

	for i := range splitted1 {
		if !strings.Contains(splitted1[i], "=") {
			continue
		}
		splitted2 := strings.SplitAfter(splitted1[i], "=")

		key := strings.Replace(splitted2[0], "=", "", -1)
		key = strings.Replace(key, "&", "", -1)
		value := strings.Replace(splitted2[1], "=", "", -1)
		value = strings.Replace(value, "&", "", -1)

		urlMap[key] = value
	}
	return urlMap
}
