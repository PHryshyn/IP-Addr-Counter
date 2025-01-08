package counter

import (
	"IP-Addr-Counter/calculator"
	"IP-Addr-Counter/processor"
	"archive/zip"
	"bufio"
	"log"
)

const Ip4MaxAddressNumber uint64 = 4294967296
const NumberBitesInInt64 uint64 = 64
const StorageSize = uint(Ip4MaxAddressNumber / NumberBitesInInt64)

func ReadFileAncCalculateUniqueIds(filePath string) (uint64, error) {
	zipFile, err := zip.OpenReader(filePath)
	if err != nil {
		log.Fatal("File not found or corrupted: ", err)
		return 0, err
	}
	defer zipFile.Close()

	var result uint64 = 0

	for _, file := range zipFile.File {
		if file.UncompressedSize64 != 0 {
			log.Println("File is found: ", file.Name)

			textFile, err := file.Open()
			if err != nil {
				log.Fatal("ERROR: Error opening file in ZIP:", err)
				continue
			}
			defer textFile.Close()

			scanner := bufio.NewScanner(textFile)

			var processedLineCount uint64 = 0
			var storage = processor.NewStorage(StorageSize)

			for scanner.Scan() {
				var ip = scanner.Text()
				var ipIndex = calculator.CalculateIndex(ip)
				if ipIndex < 0 {
					log.Println("ERROR: Invalid IP address: ", ip)
					continue
				}
				processor.ProcessIndex(uint64(ipIndex), storage)
				processedLineCount++
			}

			if err := scanner.Err(); err != nil {
				log.Fatal(err)
			}

			var calculatedUniqueIPs = processor.CalculateNumberOfUniqueIds(storage)
			log.Println("INFO: Processed lines: ", processedLineCount, " in file: ", file.Name, " with unique IPs: ", calculatedUniqueIPs)
			result += calculatedUniqueIPs
		}
	}

	return result, err
}
