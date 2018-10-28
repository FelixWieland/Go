package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	fto "github.com/jlaffaye/ftp"
)

var sendKeyword string
var receiveKeyword string

//Config File structure
type Config struct {
	Servers []FTPServer      `json:"servers"`
	Paths   []DirectoryPaths `json:"paths"`
}

//FTPServer config File structure
type FTPServer struct {
	Name     string `json:"name"`
	IP       string `json:"ip"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

//DirectoryPaths config File structure
type DirectoryPaths struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

type startAndDest struct {
	name string
	path string
}

func main() {

	sendKeyword = "to"
	receiveKeyword = "from"

	configJSONFile := readConfigFile("config.json")
	defer configJSONFile.Close()

	ftpConfig := fileToJSON(configJSONFile)

	//EX: ftpPushPull "AbsolutePath" -> desktop "testFolder"
	//OR: ftpPushPull folder <- desktop "testFolder"
	osArgs := os.Args
	typeOfArgs := getTypeOfArgs(&osArgs)
	if typeOfArgs == receiveKeyword || typeOfArgs == sendKeyword {
		startAndDestArr := getStartAndDest(&osArgs, typeOfArgs)

		serverData := getServer(ftpConfig.Servers, startAndDestArr[1])
		connection := connectToFTPServer(serverData)

		pathStruct, _ := getPath(ftpConfig.Paths, startAndDestArr[0])
		pathStructFTP, err := getPath(ftpConfig.Paths, startAndDestArr[2])
		pathFTP := ""
		if err != nil {
			pathFTP = startAndDestArr[2].path
		} else {
			pathFTP = pathStructFTP.Path
		}

		ftpActions(connection, pathStruct.Path, pathFTP, typeOfArgs)

	} else {
		info := getInfo(&osArgs)
		_ = info
	}

	_ = ftpConfig

}

func getPath(paths []DirectoryPaths, pathToSearch startAndDest) (DirectoryPaths, error) {
	for _, e := range paths {
		if e.Name == pathToSearch.name {
			return e, nil
		}
	}
	return DirectoryPaths{}, errors.New("NotFoundInConfig")
}

func getServer(servers []FTPServer, destArgs startAndDest) FTPServer {
	for _, e := range servers {
		if e.Name == destArgs.name {
			return e
		}
	}
	panic("NotFoundInConfig: " + destArgs.name)
}

func getStartAndDest(osArgs *[]string, typeOfArgs string) [3]startAndDest {

	sAd := [3]startAndDest{}
	startSetted := false

	start := startAndDest{}
	dest := startAndDest{}
	last := startAndDest{}

	for i, e := range *osArgs {
		if i == 0 || e == typeOfArgs {
			if e == typeOfArgs {
				startSetted = true
			}
			continue
		}

		if !startSetted {
			if !isPath(e) {
				start.name = string(e)
			} else {
				start.path = string(e)
			}
		} else if dest.name == "" {
			if !isPath(e) {
				dest.name = string(e)
			} else {
				dest.path = string(e)
			}
		} else {
			if !isPath(e) {
				last.name = string(e)
			} else {
				last.path = string(e)
			}
		}
	}

	sAd[0] = start
	sAd[1] = dest
	sAd[2] = last

	return sAd

}

//TODO: Not fully implementet yet
func getInfo(osArgs *[]string) string {

	for i := range *osArgs {
		if i == 0 {
			continue
		}
	}
	return "None"
}

func getTypeOfArgs(osArgs *[]string) string {
	for _, e := range *osArgs {
		if e == "from" || e == "to" {
			return e
		}
	}
	return "info"
}

func connectToFTPServer(server FTPServer) *fto.ServerConn {

	client, err := fto.Dial(server.IP + ":" + server.Port)
	if err != nil {
		panic(err)
	}

	if err := client.Login(server.Username, server.Password); err != nil {
		panic(err)
	}

	return client
}

func readConfigFile(filePath string) *os.File {
	// Open our jsonFile
	jsonFile, err := os.Open(filePath)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	return jsonFile
}

func fileToJSON(pJSONFile *os.File) Config {
	// read our opened xmlFile as a byte array.
	byteValue, err := ioutil.ReadAll(pJSONFile)
	if err != nil {
		panic(err)
	}

	// we initialize our Users array
	var config Config

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &config)

	return config
}

func isPath(str string) bool {
	return strings.Contains(str, "/") || strings.Contains(str, "\\")
}

//TODO: Implement working up- and download
func ftpActions(conn *fto.ServerConn, clientPath string, serverPath string, action string) {
	if action == receiveKeyword {
		//Download

	} else { //"to"
		//Upload
		reader, _ := os.Open("config.json")
		conn.Stor(serverPath+"/config.json", reader)
		fmt.Print("test")
	}
}
