package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/pkg/errors"
)

const prompt = "go-db"

type Command int

const (
	Exit Command = iota
	Insert
	Select
	Open
	Unknown
)

var replLoop bool = true

// Repl runs a repl loop to read, parse and execute user input
func repl(errc chan<- error) {
	for replLoop {
		input, err := getInput()
		if err != nil {
			errc <- err
		}
		// If the input is empty, skip the remaining parts of the loop
		if input == "" {
			continue
		}
		command, err := getCommand(input)
		if err != nil {
			errc <- err
		}
		executor, err := getExecutor(command)
		if err != nil {
			errc <- err
		}
		result, err := executor(input)
		if err != nil {
			errc <- err
		}
		fmt.Println(result)
	}

	errc <- nil
}

func getInput() (string, error) {
	// Get the user input
	fmt.Printf("%s> ", prompt)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", errors.Wrap(err, "could not read input")
	}

	// Trim any spaces such as new line
	input = strings.TrimSpace(input)
	return input, nil
}

// Determine the command from the user input
func getCommand(input string) (Command, error) {
	output := strings.Split(input, " ")
	if len(output) == 0 {
		return Unknown, nil
	}

	switch strings.ToLower(output[0]) {
	case "exit":
		return Exit, nil
	case "insert":
		return Insert, nil
	case "select":
		return Select, nil
	case "open":
		return Open, nil
	default:
		return Unknown, nil
	}
}

// Choose an executor based on the command identified
func getExecutor(command Command) (func(input string) (string, error), error) {
	switch command {
	case Exit:
		return execExit, nil
	case Insert:
		return execInsert, nil
	case Select:
		return execSelect, nil
	case Open:
		return execOpen, nil
	}

	return defaultExecutor, nil
}

// This is the default executor used when no other executor is appropriate
func defaultExecutor(input string) (string, error) {
	return "Unknown command.", nil
}

// This executor disables the repl loop so stop further use input
func execExit(input string) (string, error) {
	replLoop = false
	return "Bye.", nil
}
