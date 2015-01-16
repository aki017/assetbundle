package utils

import (
	"bytes"
	"encoding/binary"
	"io"

	"github.com/Sirupsen/logrus"
)

// NULL is 0x00
const NULL = 0x00

// ReadCString is READING char[size]
func ReadCString(reader io.Reader) string {
	var buffer bytes.Buffer
	for {
		c := ReadByte(reader)
		if c == 0x00 {
			break
		} else {
			buffer.WriteByte(c)
		}
	}

	return buffer.String()
}

func WriteCString(writer io.Writer, str string) (int, error) {
	writer.Write([]byte(str))
	writer.Write([]byte{0})
	return len(str) + 1, nil
}

// ReadByte is read byte
func ReadByte(reader io.Reader) byte {
	b := make([]byte, 1, 1)
	_, err := reader.Read(b)
	if err != nil {
		logrus.Panic(err)
		logrus.Panic("Cannot Read Byte")
	}
	return b[0]
}

// ReadInt32 is return int32
func ReadInt32(r io.Reader, e binary.ByteOrder) (o int32) {
	binary.Read(r, e, &o)
	return
}

// WriteInt32 is write int32
func WriteInt32(w io.Writer, o int32, e binary.ByteOrder) {
	binary.Write(w, e, &o)
	return
}

// ReadFloat32 is return float32
func ReadFloat32(r io.Reader, e binary.ByteOrder) (o float32) {
	binary.Read(r, e, &o)
	return
}

// ReadInt is read int
func ReadInt(reader io.Reader) uint32 {
	b1 := ReadByte(reader)
	b2 := ReadByte(reader)
	b3 := ReadByte(reader)
	b4 := ReadByte(reader)
	return uint32(uint32(b4) | uint32(b3)<<8 | uint32(b2)<<16 | uint32(b1)<<24)
}

// Swap is Convert Endian
func Swap(v uint32) uint32 {
	b1 := v >> 24 & 0xFF
	b2 := v >> 16 & 0xFF
	b3 := v >> 8 & 0xFF
	b4 := v & 0xFF
	return uint32(uint32(b1) | uint32(b2)<<8 | uint32(b3)<<16 | uint32(b4)<<24)
}
