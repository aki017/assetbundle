package assetbundle

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"io"
	"io/ioutil"
	"os"

	"code.google.com/p/lzma"

	"github.com/Sirupsen/logrus"
	"github.com/aki017/assetbundle/body"
	"github.com/aki017/assetbundle/header"
	"github.com/aki017/assetbundle/utils"
)

func DecodeHeader(b []byte) *header.Header {
	buffer := bytes.NewBuffer(b)
	h := new(header.Header)

	magick := utils.ReadCString(buffer)
	switch magick {
	case "UnityWeb":
		h.FileType = header.UnityWeb
	case "UnityRaw":
		h.FileType = header.UnityRaw
	}

	h.Format = header.Format(utils.ReadInt32(buffer, binary.BigEndian))
	h.PlayerVersion = utils.ReadCString(buffer)
	h.EngineVersion = utils.ReadCString(buffer)
	h.FileSize = utils.ReadInt32(buffer, binary.BigEndian)
	h.DataOffset = utils.ReadInt32(buffer, binary.BigEndian)
	h.Unknown1 = utils.ReadInt32(buffer, binary.BigEndian)

	return h
}

func decodeBodies(b []byte) (bodies []*body.Body) {
	buffer := bytes.NewBuffer(b)

	bl := utils.ReadInt32(buffer, binary.BigEndian)
	bodies = make([]*body.Body, bl)
	for i := int32(0); i < bl; i++ {
		name := utils.ReadCString(buffer)
		offset := utils.ReadInt32(buffer, binary.BigEndian)
		length := utils.ReadInt32(buffer, binary.BigEndian)
		if length > 0 {
			arg := b[offset : offset+length]
			bodies[i] = body.Decode(name, arg)
		}
	}
	return
}

// Decode reads a AssetBundle from r and returns it as an assetbundle.AssetBundle
func Decode(arg io.Reader) (ab *AssetBundle) {
	original, _ := ioutil.ReadAll(arg)

	ab = new(AssetBundle)
	ab.Header = DecodeHeader(original)
	if ab.Header.FileType == header.UnityWeb {
		bodydata, _ := ioutil.ReadAll(lzma.NewReader(bytes.NewBuffer(original[ab.Header.DataOffset:])))
		ab.Bodies = decodeBodies(bodydata)
	} else {
		ab.Bodies = decodeBodies(original[ab.Header.DataOffset:])
	}
	return
}

// DecodeFile reads a AssetBundle from path and returns it as an assetbundle.AssetBundle
func DecodeFile(path string) (ab *AssetBundle) {
	fp, err := os.Open(path)
	defer fp.Close()
	if err != nil {
		logrus.Fatalf("Cannot open file (%s)", err)
	}
	r := bufio.NewReader(fp)
	ab = Decode(r)
	return
}
