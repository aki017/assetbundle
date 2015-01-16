package body

import (
	"encoding/binary"
	"io"

	"github.com/Sirupsen/logrus"
	"github.com/aki017/assetbundle/utils"
)

func encodeHeader(r io.Reader) (h *Header) {
	h = &Header{
		TreeSize:   utils.ReadInt32(r, binary.BigEndian),
		FileSize:   utils.ReadInt32(r, binary.BigEndian),
		Format:     utils.ReadInt32(r, binary.BigEndian),
		DataOffset: utils.ReadInt32(r, binary.BigEndian),
		Reserved:   utils.ReadInt32(r, binary.BigEndian),
	}
	return
}

func encodeTypeTree(r io.Reader) (t *TypeTree) {
	t = &TypeTree{
		UnityVersion: utils.ReadCString(r),
		Version:      utils.ReadInt32(r, binary.LittleEndian),
	}

	count := utils.Swap(utils.ReadInt(r))
	t.Fields = make([]*TypeField, count, count)

	for i := uint32(0); i < count; i++ {
		t.Fields[i] = decodeTypeField(ClassID(utils.Swap(utils.ReadInt(r))), r)
	}

	_ = utils.ReadInt(r)

	return
}

func encodeTypeField(id ClassID, r io.Reader) (t *TypeField) {
	t = &TypeField{
		ID:        id,
		Type:      utils.ReadCString(r),
		Name:      utils.ReadCString(r),
		Size:      utils.ReadInt32(r, binary.LittleEndian),
		Index:     utils.ReadInt32(r, binary.LittleEndian),
		ArrayFlag: utils.ReadInt32(r, binary.LittleEndian),
		Flags1:    utils.ReadInt32(r, binary.LittleEndian),
		Flags2:    utils.ReadInt32(r, binary.LittleEndian),
	}

	count := utils.Swap(utils.ReadInt(r))

	for i := uint32(0); i < count; i++ {
		t.Children = append(t.Children, decodeTypeField(id, r))
	}

	return
}

func encodeObjects(r io.Reader) (o *Objects) {
	count := utils.ReadInt32(r, binary.LittleEndian)
	o = &Objects{
		List: make([]*Object, count, count),
	}

	for i := int32(0); i < count; i++ {
		o.List[i] = decodeObject(r)
	}

	return
}

func encodeObject(r io.Reader) (o *Object) {
	o = &Object{
		ID:       utils.ReadInt32(r, binary.LittleEndian),
		Offset:   utils.ReadInt32(r, binary.LittleEndian),
		Length:   utils.ReadInt32(r, binary.LittleEndian),
		ClassID1: ClassID(utils.ReadInt32(r, binary.LittleEndian)),
		ClassID2: ClassID(utils.ReadInt32(r, binary.LittleEndian)),
	}
	return
}
func encodeAssetRefs(r io.Reader) (a *AssetRefs) {
	count := utils.ReadInt32(r, binary.LittleEndian)

	a = &AssetRefs{
		Unknown: utils.ReadByte(r),
		List:    make([]*AssetRef, count, count),
	}

	for i := int32(0); i < count; i++ {
		a.List[i] = decodeAssetRef(r)
	}

	return
}

func encodeAssetRef(r io.Reader) (a *AssetRef) {
	a = &AssetRef{
		GUID:      NewGUID(r),
		Type:      utils.ReadInt32(r, binary.BigEndian),
		FilePath:  utils.ReadCString(r),
		AssetPath: utils.ReadCString(r),
	}
	return
}

// Encode byte array
func (body *Body) Encode(w io.Writer) {
	if body.raw != nil {
		w.Write(body.raw)
	} else {
		logrus.Panic("Unimplemented")
	}
}
