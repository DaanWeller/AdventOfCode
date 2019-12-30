package main

import (
	"fmt"
)

func main() {
	args := os.Args[1:]

	filename := args[0]
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}


