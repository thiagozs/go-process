package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	go processSignal()
	done := make(chan bool, 1)
	fmt.Println("Open a new terminal, get the PID and issue kill -# PID command")

	<-done
}

func processSignal() {

	sigch := make(chan os.Signal, 1)

	// MATCH - signal cannot be trapped
	// os.Kill
	// syscall.SIGKILL
	// syscall.SIGSTOP

	signal.Notify(sigch,
		syscall.SIGINT, syscall.SIGQUIT,
		syscall.SIGTERM, syscall.SIGUSR1,
		syscall.SIGUSR2, syscall.SIGHUP,
		os.Interrupt)

	for {
		signalType := <-sigch
		fmt.Println(">>> Received signal from channel : ", signalType)

		switch signalType {
		case syscall.SIGHUP:
			fmt.Println("+ Hangup/SIGHUP - portable number 1")
		case syscall.SIGINT:
			fmt.Println("+ Terminal interrupt signal/SIGINT - portable number 2")
		case syscall.SIGQUIT:
			fmt.Println("+ Terminal quit signal/SIGQUIT - portable number 3 - will core dump")
		case syscall.SIGABRT:
			fmt.Println("+ Process abort signal/SIGABRT - portable number 6 - will core dump")
		//case syscall.SIGKILL:
		//	fmt.Println("+ Kill signal/SIGKILL - portable number 9")
		case syscall.SIGALRM:
			fmt.Println("+ Alarm clock signal/SIGALRM - portable number 14")
		case syscall.SIGTERM:
			fmt.Println("+ Termination signal/SIGTERM - portable number 15")
		case syscall.SIGUSR1:
			fmt.Println("+ User-defined signal 1/SIGUSR1")
			//HINT : this is where you want to tell your program
			//do something upon receiving a user-defined-signal
		case syscall.SIGUSR2:
			fmt.Println("+ User-defined signal 2/SIGUSR2")
		default:
			fmt.Printf("+ signal = %v\n", signalType)
		}
	}
}
