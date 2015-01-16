package assetbundle

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"io"

	"code.google.com/p/lzma"

	"github.com/aki017/assetbundle/body"
	"github.com/aki017/assetbundle/header"
	"github.com/aki017/assetbundle/utils"
)

func encodeBodies(b []byte) (bodies []*body.Body) {
	buffer := bytes.NewBuffer(b)

	bl := utils.ReadInt32(buffer, binary.BigEndian)

	bodies = make([]*body.Body, bl)
	for i := int32(0); i < bl; i++ {
		name := utils.ReadCString(buffer)
		offset := utils.ReadInt(buffer)
		length := utils.ReadInt(buffer)
		if length > 0 {
			arg := b[offset : offset+length]
			bodies[i] = body.Decode(name, arg)
		}
	}
	return
}

// Encode writes a AssetBundle from r and returns it as an assetbundle.AssetBundle
func (ab *AssetBundle) Encode(w io.Writer) {
	var bodyBuffer bytes.Buffer
	var bodyWriter io.Writer
	switch ab.Header.FileType {
	case header.UnityWeb:
		bodyWriter = lzma.NewWriter(&bodyBuffer)
	case header.UnityRaw:
		bodyWriter = bufio.NewWriter(&bodyBuffer)
	}

	offset := 0
	utils.WriteInt32(bodyWriter, int32(len(ab.Bodies)), binary.BigEndian)
	var rawbuffer = make([]bytes.Buffer, len(ab.Bodies))
	headerSize := 4
	for _, body := range ab.Bodies {
		headerSize += len(body.Name) + 1
		headerSize += 4
		headerSize += 4
	}

	for i, body := range ab.Bodies {
		utils.WriteCString(bodyWriter, body.Name)
		body.Encode(bufio.NewWriter(&rawbuffer[i]))
		utils.WriteInt32(bodyWriter, int32(headerSize+offset), binary.BigEndian)
		utils.WriteInt32(bodyWriter, int32(rawbuffer[i].Len()), binary.BigEndian)
		offset += rawbuffer[i].Len()
	}

	for _, b := range rawbuffer {
		b.WriteTo(bodyWriter)
	}

	if ab.Header.FileType == header.UnityWeb {
		bodyWriter.(io.WriteCloser).Close()
	}

	ab.Header.FileSize = int32(bodyBuffer.Len())
	ab.Header.EncodeHeader(w)

	bodyBuffer.WriteTo(w)
}
