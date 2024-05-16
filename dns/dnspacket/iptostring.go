package dnspacket

import (
	"net"
)

func IptoString(repsonse DNSPacket) string {

	ip := repsonse.Answers[0].Data
	return net.IP(ip).String()

}
