package dnsresolver

import (
	"DNS/constants"
	"DNS/dnspacket"
	"fmt"
)

func Resolve(domainName string, recordType uint16) (string, error) {
	nameserver := "198.41.0.4"
	for {
		fmt.Printf("Querying %s for %s\n", nameserver, domainName)
		response, err := dnspacket.SendQuery(nameserver, domainName, recordType)
		if err != nil {
			return "", err
		}
		if ip := GetAnswer(*response); ip != "" {
			return ip, nil
		} else if nsIP := GetNameserverIP(*response, recordType); nsIP != "" {
			nameserver = nsIP
		} else if nsDomain := GetNameserver(*response); nsDomain != "" {
			ip, err := Resolve(nsDomain, uint16(constants.TYPE_A))
			if err != nil {
				return "", err
			}
			nameserver = ip
		} else {
			return "", fmt.Errorf("something went wrong")
		}
	}
}
