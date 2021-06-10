package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// https://stackoverflow.com/questions/28081486/how-can-i-go-run-a-project-with-multiple-files-in-the-main-package

func main() {

	errc := make(chan error)

	go run(errc)

	if err := <-errc; err != nil {
		fmt.Fprintf(os.Stderr, "FATAL: %v\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}

func run(errc chan<- error) {
	defer close(errc)

	go repl(errc)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop
	fmt.Fprintln(os.Stdout)
}
