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

func Repl() error {
	for {
		input, err := getInput()
		if err != nil {
			return err
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
}

func getInput() (string, error) {
	fmt.Printf("%s>", prompt)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		return "", errors.Wrap(err, "could not read input")
	}

	return input, nil
}

func getCommand(input string) (Command, error) {
	output := strings.Split(input, " ")
	if len(output) == 0 {
		return Unknown, errors.New("No command detected.")
	}

	switch strings.ToLower(output[0]) {
	case "exit":
		return Exit, nil
	case "write":
		return Write, nil
	case "read":
		return Read, nil
	default:
		return Unknown, errors.New("Unknown command.")
	}
}

func getExecutor(command Command) (func(input string) (string, error), error) {
	switch command {
	case Exit:
		return exitExecutor, nil
	}

	return testExecutor, nil
}

func exitExecutor(input string) (string, error) {
	panic("ahhhh")
}

func testExecutor(input string) (string, error) {
	return input, nil
}
