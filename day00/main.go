package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

var input = flag.String("input", "", "pass the filename to the input")

func main() {
	flag.Parse()
	file, err := os.Open(*input)
	if err != nil {
		log.Fatalf("%v - for file %s", err, *input)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
