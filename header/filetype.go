package header

const (
	// UnityWeb is Compressed file
	UnityWeb = FileType(iota)
	// UnityRaw is Raw file
	UnityRaw
)

// FileType is Enum
type FileType byte

func (t FileType) String() string {
	switch t {
	case UnityRaw:
		return "Compressed"
	case UnityWeb:
		return "Raw"
	}
	return "UnKnown"
}
