package main

import (
	"DNS/constants"
	"DNS/dnsresolver"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

var RecordTypes map[string]constants.RecordType = map[string]constants.RecordType{
	"A":     constants.TYPE_A,
	"NS":    constants.TYPE_NS,
	"CNAME": constants.TYPE_CNAME,
	"TXT":   constants.TYPE_TXT,
	"AAAA":  constants.TYPE_AAAA,
}

var reverseRecordTypes = map[int]string{
	1:  "A",
	2:  "NS",
	5:  "CNAME",
	16: "TXT",
	28: "AAAA",
}

func resolve(name string, t constants.RecordType) []string {
	// most of your code should go here. use a switch statement
	// so each resolution type goes into a different function

	filePath := "dns_cache.json"

	dnsCache := dnsresolver.NewDNSCache(filePath)
	dnsCache.LoadFromDisk()
	resolvedValue := make([]string, 0, 0)

	ip := dnsCache.Get(name, reverseRecordTypes[int(t)])
	if ip != "" {
		resolvedValue = append(resolvedValue, ip)
		// fmt.Printf("Found IP: %s", ip)

	} else {

		IP, err := dnsresolver.Resolve(name, uint16(t))
		if err != nil {
			fmt.Println("Failed to send Query:", err)
			return nil
		}
		resolvedValue = append(resolvedValue, IP)
		dnsCache.Add(name, reverseRecordTypes[int(t)], IP, 60*time.Second)

	}

	return resolvedValue
}

func main() {
	// get all command line arguments
	names := os.Args[1:]
	t := flag.String("t", "A", "the record type to query for each name")
	flag.Parse()

	// input validation
	if len(names) == 0 {
		fmt.Println("Not enough arguments, must pass in at least one name")
		os.Exit(1)
	}

	if _, exists := RecordTypes[*t]; !exists {
		keys := make([]string, 0, len(RecordTypes))
		for k := range RecordTypes {
			keys = append(keys, k)
		}
		fmt.Printf("Specified record type %s doesn't exist. Must be one of %v", *t, keys)
		os.Exit(1)
	}
	names = flag.Args()
	// invoke the resolve function for each of the given names
	for _, name := range names {
		fmt.Printf("%s,%s", name, strings.Join(resolve(name, RecordTypes[*t]), ""))
		println("")
	}
}
