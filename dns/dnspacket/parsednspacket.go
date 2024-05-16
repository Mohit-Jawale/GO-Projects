package dnspacket

import (
	"DNS/dnsrequest"
	"DNS/dnsresponse"
	"bytes"
	"fmt"
)

func ParseDNSPacket(data []byte) (*DNSPacket, error) {
	reader := bytes.NewReader(data)

	header, err := dnsresponse.ParseHeader(reader)
	if err != nil {
		return nil, fmt.Errorf("error parsing header: %w", err)
	}

	questions := make([]dnsrequest.DNSQuestion, header.NumQuestions)
	for i := range questions {
		question, err := dnsresponse.ParseQuestion(reader)
		if err != nil {
			return nil, fmt.Errorf("error parsing question: %w", err)
		}
		questions[i] = *question
	}

	answers := make([]dnsresponse.DNSRecord, header.NumAnswers)
	for i := range answers {
		answer, err := dnsresponse.ParseRecord(reader)
		if err != nil {
			return nil, fmt.Errorf("error parsing answer: %w", err)
		}
		answers[i] = *answer
	}

	authorities := make([]dnsresponse.DNSRecord, header.NumAuthorities)
	for i := range authorities {
		authority, err := dnsresponse.ParseRecord(reader)
		if err != nil {
			return nil, fmt.Errorf("error parsing authority: %w", err)
		}
		authorities[i] = *authority
	}

	additionals := make([]dnsresponse.DNSRecord, header.NumAdditionals)
	for i := range additionals {
		additional, err := dnsresponse.ParseRecord(reader)
		if err != nil {
			return nil, fmt.Errorf("error parsing additional: %w", err)
		}
		additionals[i] = *additional
	}

	return &DNSPacket{
		Header:      *header,
		Questions:   questions,
		Answers:     answers,
		Authorities: authorities,
		Additionals: additionals,
	}, nil
}
