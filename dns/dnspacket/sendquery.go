package dnspacket

import (
	"DNS/constants"
	"DNS/dnsquerybuilder"
	"fmt"
	"net"
)

func SendQuery(ipaddress string, domainName string, recordType uint16) (*DNSPacket, error) {

	query, err := dnsquerybuilder.BuildQuery(domainName, recordType)

	if recordType == uint16(constants.TYPE_AAAA) {

		ipaddress = "[" + ipaddress + "]"

	}
	if err != nil {
		panic(err)
	}

	// Create a udp socket
	conn, err := net.Dial("udp", ipaddress+":53")

	if err != nil {
		fmt.Println("Failed to create socket:", err)
		return nil, nil
	}
	defer conn.Close()

	_, err = conn.Write(query)
	if err != nil {
		fmt.Println("Failed to send query:", err)
		return nil, nil
	}

	buffer := make([]byte, 1024)

	// Read the response from the DNS server
	_, err = conn.Read(buffer)
	if err != nil {
		fmt.Println("Failed to read response:", err)
		return nil, nil
	}

	packet, err := ParseDNSPacket(buffer)

	return packet, err

}
