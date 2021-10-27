package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/go-chrono/chrono"
)

func main() {
	parseCmd := flag.NewFlagSet("parse", flag.ExitOnError)

	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "parse":
		argument, err := getFirstArgument()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		parseCmd.Parse([]string{argument})

		duration, err := parse(argument)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		fmt.Println(getHourMinuteString(duration))
	default:
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Usage: duration [command] [argument]")
	flag.PrintDefaults()
}

func getFirstArgument() (string, error) {
	if isInputFromPipe() {
		bytes, err := io.ReadAll(os.Stdin)
		return strings.TrimSpace(string(bytes)), err
	}

	argument := flag.Arg(0)
	if len(argument) == 0 {
		return "", errors.New("Missing argument")
	}

	return argument, nil
}

func isInputFromPipe() bool {
	fileInfo, _ := os.Stdin.Stat()
	return fileInfo.Mode()&os.ModeCharDevice == 0
}

func parse(durationStr string) (chrono.Duration, error) {
	var d chrono.Duration
	err := d.Parse(durationStr)
	return d, err
}

func getHourMinuteString(duration chrono.Duration) string {
	hours := int(duration.Seconds() / 60 / 60)
	minutes := int(duration.Seconds()/60) % 60
	seconds := int(duration.Seconds()) % 60
	return intToStr(hours) + ":" + intToStr(minutes) + ":" + intToStr(seconds)
}

func intToStr(n int) string {
	str := strconv.Itoa(n)
	if len(str) == 1 {
		str = "0" + str
	}
	return str
}
