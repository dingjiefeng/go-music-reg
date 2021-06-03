package audio

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func bits64ToInt(b []byte) uint64 {
	if len(b) != 8 {
		panic("Expected size 8!")
	}
	var payload uint64
	buf := bytes.NewReader(b)
	err := binary.Read(buf, binary.LittleEndian, &payload)
	if err != nil {
		panic(err)
	}
	return payload
}

func bits32ToInt(b []byte) int {
	if len(b) != 4 {
		panic("Expected size 4!")
	}
	var payload int32
	buf := bytes.NewReader(b)
	err := binary.Read(buf, binary.LittleEndian, &payload)
	if err != nil {
		panic(err)
	}
	return int(payload) // easier to work with ints
}

func bits16ToInt(b []byte) int {
	if len(b) != 2 {
		fmt.Println(len(b))
		panic("Expected size 2!")
	}
	var payload int16
	buf := bytes.NewReader(b)
	err := binary.Read(buf, binary.LittleEndian, &payload)
	if err != nil {
		// TODO: make safe
		panic(err)
	}
	return int(payload) // easier to work with ints
}
