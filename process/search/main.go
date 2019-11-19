package main

import (
	"flag"
	ps "github.com/mitchellh/go-ps"
	"log"
	"strings"
)

func main() {
	// flag names
	exec := flag.String("exec", "foo", "executable name")

	// get all process
	processList, err := ps.Processes()
	if err != nil {
		log.Println("ps.Processes() Failed, are you using windows?")
		return
	}

	// parse flags
	flag.Parse()

	// if flag exec is empty change this
	if *exec == "" || *exec == "foo" {
		*exec = "main"
	}

	// find the flag with the name
	for x := range processList {
		var process ps.Process = processList[x]
		if strings.Contains(process.Executable(), *exec) {
			log.Printf("%d\t%d\t%s\n", process.Pid(), process.PPid(), process.Executable())
		}
	}

}
