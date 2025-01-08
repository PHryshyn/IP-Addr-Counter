package main

import (
	"IP-Addr-Counter/counter"
	"log"
)

func main() {
	var fileName = "/Users/pavelhryshyn/Downloads/"

	var result, err = counter.ReadFileAncCalculateUniqueIds(fileName)

	if err != nil {
		log.Fatal("ERROR: Error during processing:", err)
	}

	log.Println("INFO: Total unique IPs: ", result)
}
