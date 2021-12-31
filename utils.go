package main

import (
	"bytes"
	"encoding/binary"
	"log"
)

//converts int64 to byte array
func IntToHex(integer int64) []byte {
	buffer := new(bytes.Buffer)
	erro := binary.Write(buffer, binary.BigEndian, integer)
	if erro != nil {
		log.Panic(erro)
	}
	return buffer.Bytes()
}