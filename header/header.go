package header

// Header is AssetBundle Header
type Header struct {
	FileType      FileType
	Format        Format
	PlayerVersion string
	EngineVersion string
	FileSize      int32
	DataOffset    int32
	Unknown1      int32
}

// PlayerVersion is String
type PlayerVersion string

// EngineVersion is String
type EngineVersion string
