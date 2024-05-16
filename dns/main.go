package main

import (
	"DNS/constants"
	"DNS/dnsresolver"
	"fmt"
)

// import (
// 	"flag"
// 	"fmt"
// 	"os"
// 	"strings"
// )

type RecordType uint16

// var RecordTypes map[string]RecordType = map[string]RecordType{
// 	"A":     TYPE_A,
// 	"NS":    TYPE_NS,
// 	"CNAME": TYPE_CNAME,
// 	"TXT":   TYPE_TXT,
// 	"AAAA":  TYPE_AAAA,
// }

// func resolve(name string, t RecordType) []string {
// 	// most of your code should go here. use a switch statement
// 	// so each resolution type goes into a different function
// 	resolvedValue := make([]string, 0, 0)
// 	return resolvedValue
// }

// func main() {
// 	// get all command line arguments
// 	names := os.Args[1:]
// 	t := flag.String("type", "A", "the record type to query for each name")
// 	flag.Parse()

// 	// input validation
// 	if len(names) == 0 {
// 		fmt.Println("Not enough arguments, must pass in at least one name")
// 		os.Exit(1)
// 	}

// 	if _, exists := RecordTypes[*t]; !exists {
// 		keys := make([]string, 0, len(RecordTypes))
// 		for k := range RecordTypes {
// 			keys = append(keys, k)
// 		}
// 		fmt.Printf("Specified record type %s doesn't exist. Must be one of %v", *t, keys)
// 		os.Exit(1)
// 	}

// 	// invoke the resolve function for each of the given names
// 	for _, name := range names {
// 		fmt.Printf("%s,%s", name, strings.Join(resolve(name, RecordTypes[*t]), ""))
// 	}
// }

func main() {

	IP, err := dnsresolver.Resolve("x.com", uint16(constants.TYPE_A))
	if err != nil {
		fmt.Println("Failed to send Query:", err)
		return
	}
	fmt.Println(IP)

}
