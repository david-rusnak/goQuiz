package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	data := readCsv("./test.csv")

	for i:=0; i < len(data); i++ {
		fmt.Println(data[i][0])
		var answer string
		fmt.Scanln(&answer)
		
		if answer == data[i][1] {
			fmt.Println("u go girl")
		} else {
			log.Fatal("u suck")
		}
		fmt.Println()
	}
}

func readCsv(path string) [][]string {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal("Failed to open " + path, err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Unable to read", err)
	}

	return records
}