package main

import (
	"log"
	"os"

	"github.com/jghiloni/adventofcode2020"
)

func main() {
	if err := adventofcode2020.Main(os.Args); err != nil {
		log.Fatal(err)
	}
}
