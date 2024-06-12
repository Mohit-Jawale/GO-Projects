package dnsresolver

import (
	"DNS/constants"
	"DNS/dnspacket"
	"fmt"
	"net"
	"regexp"
	"strings"
)

func GetAnswer(records dnspacket.DNSPacket) string {
	for _, record := range records.Answers {

		if record.Type == uint16(constants.TYPE_CNAME) {
			fmt.Println(record)
			// Convert byte slice to string
			var sb strings.Builder
			for _, b := range record.Data {
				if b >= 32 && b <= 126 {
					// Printable ASCII characters
					sb.WriteByte(b)
				} else {
					// Non-printable characters
					sb.WriteString(fmt.Sprintf("\\x%02x", b))
				}
			}
			// re := regexp.MustCompile(`[^\x20-\x7E]`)
			replacedString := strings.ReplaceAll(sb.String(), "\x04", ".")
			re := regexp.MustCompile(`[^\x20-\x7E.]`)

			cleanedString := re.ReplaceAllString(replacedString, "")
			fmt.Println("Converted string:", cleanedString)
		}
		return net.IP(record.Data).String()

	}
	return ""
}
