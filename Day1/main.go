package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]

	fileName := args[0]
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	total := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Sprintf("Module mass = %s", line)
		moduleMass, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		fuel := calculateFuel(moduleMass, 0)
		total += fuel
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(total)
//	moduleMass, err := strconv.Atoi(args[0])
//	if err != nil {
//		fmt.Println("bruh")
//		panic(err)
//	}
//
//	fmt.Sprintf("Module mass = %d", moduleMass)
//	fuel := calculateFuel(moduleMass)
//	fmt.Println(fuel)
}

func calculateFuel(fuel int, result int) int {
	fuelNeeded := (fuel /3) - 2
	if fuelNeeded <= 0 {
		return result
	} else {
		result += fuelNeeded
		return calculateFuel(fuelNeeded, result)
	}
}
