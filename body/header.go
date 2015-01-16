package body

// Header is AssetBundle Body Header
type Header struct {
	TreeSize   int32
	FileSize   int32
	Format     int32
	DataOffset int32
	Reserved   int32
}
