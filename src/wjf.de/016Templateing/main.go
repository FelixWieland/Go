package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"
	"time"
)

type demo struct {
	Title string //Be careful! Go is really really Case-Sensitive
	Text  string
	Rest  []string
	Time  time.Time
}

var fm = template.FuncMap{
	"uc":     demoFunc,
	"ft":     strings.ToUpper,
	"Format": time.Now().Format,
}

func demoFunc() string {
	return "test"
}

var tpl *template.Template //package.Type

func init() {
	//Binds functions to templates
	tpl = template.Must(template.New("").Funcs(fm).ParseGlob("templates/*.html"))
}

func main() {
	sValues := demo{
		"Titel",
		"01-02-06",
		[]string{"test", "test", "test"},
		time.Now(),
	}

	//Anonymouse struct Datatype
	sDemo := struct {
		name  string
		lname string
	}{
		"James",
		"Bond",
	}

	fmt.Printf(sDemo.name)
	//err := tpl.Execute(os.Stdout, sValues)
	err := tpl.ExecuteTemplate(os.Stdout, "one.html", sValues)
	if err != nil {
		log.Fatalln(err)
	}
}