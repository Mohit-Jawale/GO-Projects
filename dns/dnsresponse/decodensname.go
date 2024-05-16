package dnsresponse

func DecodeNSName(data []byte) string {
	var name string
	for i := 0; i < len(data); {
		length := int(data[i])
		i++
		if length == 0 {
			break
		}
		if len(name) > 0 {
			name += "."
		}
		name += string(data[i : i+length])
		i += length
	}
	return name
}
