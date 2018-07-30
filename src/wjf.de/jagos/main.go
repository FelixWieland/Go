package main

import (
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"
	"time"

	"github.com/julienschmidt/httprouter"

	"wjf.de/jagos/lib"
)

var tpl *template.Template //package.Type

var templateFunctions = template.FuncMap{
	"ft":     strings.ToUpper,
	"Format": time.Now().Format,
}

func loadTemplates() {
	tpl = template.Must(template.New("").Funcs(templateFunctions).ParseGlob("templates/*.html"))
}

func init() {
	loadTemplates()
}

func main() {

	router := httprouter.New()

	//Resource routing
	router.GET("/resources/*path", resourcesServe)
	//Content delivery file routing
	router.GET("/cdf/plc/*path", plcServe)

	router.POST("/AJAX/SYSTEM/:func", systemAjax)

	if !checkIfConfigsSetuped() {
		router.GET("/", firstStart)
		router.GET("/ajax", firstStart)
	} else {
		//load config

		//Site routing
		router.GET("/", index)
	}
	//router.GET("/pvt/*path", pvtFileServe)
	http.ListenAndServe(":80", router)
}

func firstStart(w http.ResponseWriter, r *http.Request, rt httprouter.Params) {

	loadTemplates()

	err := tpl.ExecuteTemplate(w, "firstStart.html", nil)
	if err != nil {
		log.Fatalln(err)
	}
	return
}

func backend(w http.ResponseWriter, r *http.Request, rt httprouter.Params) {
	return
}

func checkIfConfigsSetuped() bool {
	if _, err := os.Stat("config/mysql.conf"); os.IsNotExist(err) {
		return false
	}
	return true
}

func index(w http.ResponseWriter, r *http.Request, rt httprouter.Params) {
	http.ServeFile(w, r, "index.html")
}

func resourcesServe(w http.ResponseWriter, r *http.Request, rt httprouter.Params) {
	http.ServeFile(w, r, "resources"+rt.ByName("path"))
}

func plcServe(w http.ResponseWriter, r *http.Request, rt httprouter.Params) {
	http.ServeFile(w, r, "cdf/plc"+rt.ByName("path"))
}

func systemAjax(w http.ResponseWriter, r *http.Request, rt httprouter.Params) {
	funcToCall := rt.ByName("func")
	switch funcToCall {
	case "setConfig":
		lib.SetConfig("", "", "")
	}
}
