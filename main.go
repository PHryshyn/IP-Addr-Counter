package main

import (
	"IP-Addr-Counter/counter"
	"log"
	"path/filepath"
)

func main() {

	filePath := filepath.Join("PLEASE Enter Path to Folder", "ip_addresses.zip")

	log.Println("File path:", filePath)

	var result, err = counter.ReadFileAncCalculateUniqueIds(filePath)

	if err != nil {
		log.Fatal("ERROR: Error during processing:", err)
	}

	log.Println("INFO: Total unique IPs: ", result)
}
