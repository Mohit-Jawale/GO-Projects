package dnsresolver

import (
	"DNS/constants"
	"DNS/dnspacket"
	"fmt"
)

func Resolve(domainName string, recordType uint16) (string, string, error) {

	nameserver := "198.41.0.4"
	var nsName string
	var nsIP string

	for {
		response, err := dnspacket.SendQuery(nameserver, domainName, recordType)
		//fmt.Println("this is nfrkjnvskjrn", response)
		if err != nil {
			return "", "", err
		}
		if ip := GetAnswer(*response); ip != "" {

			return ip, nsName, nil
		} else if nsIP, nsName = GetNameserverIP(*response, recordType); nsIP != "" {
			nameserver = nsIP
			if recordType == uint16(constants.TYPE_NS) || recordType == uint16(constants.TYPE_CNAME) {
				recordType = uint16(constants.TYPE_A)
			}
		} else if nsDomain := GetNameserver(*response); nsDomain != "" {
			ip, _, err := Resolve(nsDomain, uint16(constants.TYPE_A))

			if err != nil {
				return "", "", err
			}
			nameserver = ip
			if recordType == uint16(constants.TYPE_NS) || recordType == uint16(constants.TYPE_CNAME) {
				recordType = uint16(constants.TYPE_A)
			}

		} else {
			return "", "", fmt.Errorf("something went wrong")
		}
	}
}
