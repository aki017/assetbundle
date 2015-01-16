package body

import "hash/crc32"

// Body is body
type Body struct {
	Name      string
	Header    *Header
	TypeTree  *TypeTree
	Objects   *Objects
	AssetRefs *AssetRefs
	raw       []byte
}

// CRC is crc32
func (b Body) CRC() uint32 {
	return crc32.ChecksumIEEE(b.raw)
}

func (b *Body) SetBody(r []byte) {
	b.raw = r
}

func (b *Body) GetBody() []byte {
	return b.raw
}
