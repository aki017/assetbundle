package header

import "github.com/dustin/go-humanize"

// FileSize is filesize
type FileSize uint32

func (f FileSize) String() string {
	return humanize.Bytes(uint64(f))
}
