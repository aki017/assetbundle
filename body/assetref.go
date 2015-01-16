package body

import (
	"fmt"
	"io"

	"github.com/aki017/assetbundle/utils"
)

type AssetRefs struct {
	Unknown byte
	List    []*AssetRef
}

type AssetRef struct {
	GUID      *GUID
	Type      int32
	FilePath  string
	AssetPath string
}

// GUID is Global Unique IDentifier
type GUID struct {
	Bytes [16]byte
}

func NewGUID(r io.Reader) *GUID {
	return &GUID{
		Bytes: [16]byte{
			utils.ReadByte(r),
			utils.ReadByte(r),
			utils.ReadByte(r),
			utils.ReadByte(r),
			utils.ReadByte(r),
			utils.ReadByte(r),
			utils.ReadByte(r),
			utils.ReadByte(r),
			utils.ReadByte(r),
			utils.ReadByte(r),
			utils.ReadByte(r),
			utils.ReadByte(r),
			utils.ReadByte(r),
			utils.ReadByte(r),
			utils.ReadByte(r),
			utils.ReadByte(r),
		},
	}
}

func (g GUID) String() string {
	return fmt.Sprintf(
		`"%04x%04x-%04x-%04x-%04x-%04x%04x%04x"`,
		uint16(g.Bytes[0])<<8+uint16(g.Bytes[1]),
		uint16(g.Bytes[2])<<8+uint16(g.Bytes[3]),
		uint16(g.Bytes[4])<<8+uint16(g.Bytes[5]),
		uint16(g.Bytes[6])<<8+uint16(g.Bytes[7]),
		uint16(g.Bytes[8])<<8+uint16(g.Bytes[9]),
		uint16(g.Bytes[10])<<8+uint16(g.Bytes[11]),
		uint16(g.Bytes[12])<<8+uint16(g.Bytes[13]),
		uint16(g.Bytes[14])<<8+uint16(g.Bytes[15]),
	)
}

func (g GUID) MarshalJSON() ([]byte, error) {
	return []byte(g.String()), nil
}
