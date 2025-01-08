package calculator

import (
	"log"
	"strconv"
	"strings"
)

const (
	MAX_OCTET_NUMBER  = 256
	FIRST_OCTET_VALUE = MAX_OCTET_NUMBER * MAX_OCTET_NUMBER
	ZERO_OCTET_VALUE  = MAX_OCTET_NUMBER * MAX_OCTET_NUMBER * MAX_OCTET_NUMBER
)

func CalculateIndex(ip string) int64 {
	var splittedIp = strings.Split(ip, ".")
	if len(splittedIp) != 4 {
		log.Println("Invalid IP address: ", ip)
		return -1
	}

	return parseOstets(splittedIp[0])*ZERO_OCTET_VALUE +
		parseOstets(splittedIp[1])*FIRST_OCTET_VALUE +
		parseOstets(splittedIp[2])*MAX_OCTET_NUMBER +
		parseOstets(splittedIp[3])
}

func parseOstets(ostet string) int64 {
	number, err := strconv.Atoi(ostet)
	if err != nil || (number < 0 || number >= MAX_OCTET_NUMBER) {
		log.Println("Invalid IP ostet: ", ostet)
		return -1
	}

	return int64(number)
}
