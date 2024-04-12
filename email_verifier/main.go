package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("domain,hasMX,hasSPF,spfRecord,hasDMARC,dmarcRecord \n")

	for scanner.Scan() {
		checkDomain(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("could not  read the input %v \n", err)
	}

}
func checkDomain(domain string) {
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	mxrecords, err := net.LookupMX(domain)

	if err != nil {
		log.Printf("error:%v", err)
	}
	if len(mxrecords) > 0 {
		hasMX = true
	}

	txtrecords, err := net.LookupTXT(domain)

	if err != nil {
		log.Printf("error:%v", err)
	}

	for _, record := range txtrecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	dmarcrecords, err := net.LookupTXT("_dmarc." + domain)

	if err != nil {
		log.Printf("error:%v", err)
	}

	for _, record := range dmarcrecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}

	fmt.Printf("%v,%v,%v,%v,%v ", hasMX, hasSPF, hasDMARC, spfRecord, dmarcRecord)

}
