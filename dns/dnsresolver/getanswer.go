package dnsresolver

import (
	"DNS/constants"
	"DNS/dnspacket"
	"net"
)

func decodeCompressedDNSName(data []byte, offset int, originalData []byte) string {
	var dnsName string
	i := offset

	for i < len(data) {
		length := int(data[i])
		if length == 0 {
			break
		}
		if length&0xC0 == 0xC0 {
			// Pointer to another part of the message
			if i+1 >= len(data) {
				return dnsName // Avoid out-of-bounds error
			}
			pointer := int(data[i]&0x3F)<<8 | int(data[i+1])
			if pointer >= len(originalData) {
				return dnsName // Avoid out-of-bounds error
			}
			dnsName += decodeCompressedDNSName(originalData, pointer, originalData)
			break
		} else {
			// Read label
			i++
			if i+length > len(data) {
				return dnsName // Avoid out-of-bounds error
			}
			label := string(data[i : i+length])
			if len(dnsName) > 0 {
				dnsName += "."
			}
			dnsName += label
			i += length
		}
	}
	return dnsName
}

func GetAnswer(records dnspacket.DNSPacket) string {
	for _, record := range records.Answers {

		if record.Type == uint16(constants.TYPE_CNAME) {

			cname := decodeCompressedDNSName(record.Data, 0, record.Data)
			return cname + "." + string(record.Data)
		}

		return net.IP(record.Data).String()

	}
	return ""
}
