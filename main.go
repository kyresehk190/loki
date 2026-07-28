package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// processLine simulates the processing logic for log entries
func processLine(line string) {
	if strings.TrimSpace(line) != "" {
		fmt.Println(line)
	}
}

// readAndProcess simulates the logic that should be in logcli to handle streaming input
func readAndProcess(reader io.Reader) error {
	bufReader := bufio.NewReader(reader)
	for {
		line, err := bufReader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				if len(line) > 0 {
					processLine(line)
				}
				return nil
			}
			return err
		}
		processLine(line)
	}
}

func main() {
	// Example usage demonstrating the fix
	input := `{"log":"entry1"}
{"log":"entry2"}`
	reader := strings.NewReader(input)
	_ = readAndProcess(reader)
}