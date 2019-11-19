package main

import (
	"flag"
	"log"
	"os/exec"
	"strings"
)

func main() {
	flag.Parse()
	arg1, arg2 := strings.Join(flag.Args()[:1], " "), strings.Join(flag.Args()[1:], " ")

	tail := flag.Args()
	log.Printf("Tail: %+q\n", tail)

	cmd := exec.Command(arg1, arg2)

	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish...")
	log.Printf("Process id is %v", cmd.Process.Pid)

	err = cmd.Wait()
	if err != nil {
		log.Printf("Command finished with error, now restarting: %v", err)
		return
	}

	log.Printf("done")
}
