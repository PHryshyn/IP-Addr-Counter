package main

import (
	"IP-Addr-Counter/counter"
	"log"
)

func main() {
	var fileName = "Enter the path to the file here"

	var result, err = counter.ReadFileAncCalculateUniqueIds(fileName)

	if err != nil {
		log.Fatal("ERROR: Error during processing:", err)
	}

	log.Println("INFO: Total unique IPs: ", result)
}
