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
	Write
	Read
	Unknown
)

var replLoop bool = true

func Repl() error {
	for replLoop {
		input, err := getInput()
		if err != nil {
			return err
		}
		if input == "" {
			continue
		}
		command, err := getCommand(input)
		if err != nil {
			return err
		}
		executor, err := getExecutor(command)
		if err != nil {
			return err
		}
		result, err := executor(input)
		if err != nil {
			return err
		}
		fmt.Println(result)
	}

	return nil
}

func getInput() (string, error) {
	fmt.Printf("%s> ", prompt)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		return "", errors.Wrap(err, "could not read input")
	}

	input = strings.TrimSpace(input)
	return input, nil
}

func getCommand(input string) (Command, error) {
	output := strings.Split(input, " ")
	if len(output) == 0 {
		return Unknown, nil
	}

	switch strings.ToLower(output[0]) {
	case "exit":
		return Exit, nil
	case "write":
		return Write, nil
	case "read":
		return Read, nil
	default:
		return Unknown, nil
	}
}

func getExecutor(command Command) (func(input string) (string, error), error) {
	switch command {
	case Exit:
		return exitExecutor, nil
	}

	return defaultExecutor, nil
}

func defaultExecutor(input string) (string, error) {
	return "Unknown command.", nil
}

func exitExecutor(input string) (string, error) {
	replLoop = false
	return "Bye.", nil
}
