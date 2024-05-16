package dnsresponse

import (
	"bytes"
	"io"
)

func DecodeName(reader *bytes.Reader) ([]byte, error) {
	var parts [][]byte
	for {
		lengthByte, err := reader.ReadByte()
		if err != nil {
			return nil, err // Handle error (likely EOF)
		}
		length := int(lengthByte)

		if length == 0 {
			break // End of the name part
		}

		if length&0xC0 == 0xC0 { // Check if it's a pointer (compression)
			// Read the next byte to complete the pointer
			nextByte, err := reader.ReadByte()
			if err != nil {
				return nil, err
			}
			pointer := ((length & 0x3F) << 8) | int(nextByte)
			currentPos, _ := reader.Seek(0, io.SeekCurrent) // Store current position

			// Decode the name starting at the pointer
			result, err := DecodeCompressedName(reader, pointer)
			if err != nil {
				return nil, err
			}
			parts = append(parts, result)

			// Restore the reader's position
			reader.Seek(currentPos, io.SeekStart)
			break // After decoding a compressed name, it ends the name field
		} else { // Normal label
			part := make([]byte, length)
			_, err := reader.Read(part)
			if err != nil {
				return nil, err
			}
			parts = append(parts, part)
		}
	}

	return bytes.Join(parts, []byte(".")), nil
}

func DecodeCompressedName(reader *bytes.Reader, pointer int) ([]byte, error) {
	reader.Seek(int64(pointer), io.SeekStart)
	return DecodeName(reader)
}
