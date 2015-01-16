package header

import (
	"encoding/binary"
	"io"

	"github.com/aki017/assetbundle/utils"
)

func (h *Header) EncodeHeader(w io.Writer) {
	switch h.FileType {
	case UnityWeb:
		utils.WriteCString(w, "UnityWeb")
	case UnityRaw:
		utils.WriteCString(w, "UnityRaw")
	}
	h.DataOffset = 0x27

	utils.WriteInt32(w, int32(h.Format), binary.BigEndian)
	utils.WriteCString(w, h.PlayerVersion)
	utils.WriteCString(w, h.EngineVersion)
	utils.WriteInt32(w, h.FileSize, binary.BigEndian)
	utils.WriteInt32(w, h.DataOffset, binary.BigEndian)
	utils.WriteInt32(w, h.Unknown1, binary.BigEndian)
}
