package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	// Set the split function for the scanning operation.
	scanner.Split(bufio.ScanWords)
	// Count the words.
	params := []string{}
	for scanner.Scan() {
		params = append(params, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

	i := 1

	expectedFlagKey := ""
	var nFlag int
	var pFlag int = 1

	for ; i < len(os.Args); i++ {
		arg := os.Args[i]
		if arg[0] == '-' {
			expectedFlagKey = arg[1:]

		} else if len(expectedFlagKey) > 0 {
			switch expectedFlagKey {
			case "n":
				if nValue, err := strconv.Atoi(arg); err == nil {
					nFlag = nValue
				}
			case "P":
				if pValue, err := strconv.Atoi(arg); err == nil {
					pFlag = pValue
				}
			}
			expectedFlagKey = ""
		} else {
			break
		}
	}

	commandName := os.Args[i]
	i++

	invocationsCount := 1
	if nFlag > 0 {
		invocationsCount = len(params) / nFlag
	} else {
		nFlag = len(params)
	}

	paramsOffset := 0

	sem := make(chan int, pFlag)
	completed := make(chan bool, invocationsCount)

	for j := 0; j < invocationsCount; j++ {
		go func(paramsOffset int) {
			sem <- 1
			cmd := exec.Command(commandName, append(os.Args[i:], params[paramsOffset:paramsOffset+nFlag]...)...)
			cmdOut, err := cmd.StdoutPipe()
			if err != nil {
				log.Fatal(err)
			}

			cmdErr, err := cmd.StderrPipe()
			if err != nil {
				log.Fatal(err)
			}

			if err := cmd.Start(); err != nil {
				log.Fatal(err)
			}

			if _, err := io.Copy(os.Stdout, cmdOut); err != nil {
				log.Fatal(err)
			}
			if _, err := io.Copy(os.Stderr, cmdErr); err != nil {
				log.Fatal(err)
			}

			if err := cmd.Wait(); err != nil {
				log.Fatal(err)
			}
			<-sem
			completed <- true

		}(paramsOffset)
		paramsOffset += nFlag
	}

	for j := 0; j < invocationsCount; j++ {
		<-completed
	}

}
