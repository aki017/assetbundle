package body

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"io"
	"reflect"

	"github.com/Sirupsen/logrus"
	"github.com/aki017/assetbundle/utils"
)

func decodeHeader(r io.Reader) (h *Header) {
	h = &Header{
		TreeSize:   utils.ReadInt32(r, binary.BigEndian),
		FileSize:   utils.ReadInt32(r, binary.BigEndian),
		Format:     utils.ReadInt32(r, binary.BigEndian),
		DataOffset: utils.ReadInt32(r, binary.BigEndian),
		Reserved:   utils.ReadInt32(r, binary.BigEndian),
	}
	return
}

func decodeTypeTree(r io.Reader) (t *TypeTree) {
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

func decodeTypeField(id ClassID, r io.Reader) (t *TypeField) {
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

func decodeObjects(r io.Reader) (o *Objects) {
	count := utils.ReadInt32(r, binary.LittleEndian)
	o = &Objects{
		List: make([]*Object, count, count),
	}

	for i := int32(0); i < count; i++ {
		o.List[i] = decodeObject(r)
	}

	return
}

func decodeObject(r io.Reader) (o *Object) {
	o = &Object{
		ID:       utils.ReadInt32(r, binary.LittleEndian),
		Offset:   utils.ReadInt32(r, binary.LittleEndian),
		Length:   utils.ReadInt32(r, binary.LittleEndian),
		ClassID1: ClassID(utils.ReadInt32(r, binary.LittleEndian)),
		ClassID2: ClassID(utils.ReadInt32(r, binary.LittleEndian)),
	}
	return
}
func decodeAssetRefs(r io.Reader) (a *AssetRefs) {
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

func decodeAssetRef(r io.Reader) (a *AssetRef) {
	a = &AssetRef{
		GUID:      NewGUID(r),
		Type:      utils.ReadInt32(r, binary.BigEndian),
		FilePath:  utils.ReadCString(r),
		AssetPath: utils.ReadCString(r),
	}
	return
}

// Decode byte array
func Decode(name string, b []byte) (body *Body) {
	r := bytes.NewBuffer(b)
	body = &Body{}
	body.Name = name
	body.SetBody(b)
	body.Header = decodeHeader(r)

	switch body.Header.Format {
	case 5, 6, 7, 8:
		// TODO: Implements
		logrus.Panicf("Not Supported Format(%d)", body.Header.Format)
	case 9:
		body.TypeTree = decodeTypeTree(r)
		body.Objects = decodeObjects(r)
		body.AssetRefs = decodeAssetRefs(r)
	}

	return
}

func ParseObject(d []byte, t TypeField) UnityObject {
	r := bytes.NewBuffer(d)

	return parseObject(r, t)
}

func parseObject(r io.Reader, t TypeField) UnityObject {
	o := UnityObject{}
	o.t = &t

	for _, f := range t.Children {
		v := parseValue(r, *f)
		o.Children = append(o.Children, v)
	}

	return o
}

func wrap(t TypeField, v interface{}) UnityObject {
	return UnityObject{t: &t, data: v}

}
func parseValue(r io.Reader, t TypeField) UnityObject {
	logrus.Debugf("Start --- %s: %s(%s)", t.Name, t.Type, reflect.TypeOf(t.Type))
	switch t.Type {
	case "string":
		v := parseCollection(r, t, true)
		var b bytes.Buffer
		for _, d := range v.Children {
			b.WriteByte(d.data.(byte))
		}
		logrus.Debugf("string: %s", b.String())
		return wrap(t, b.String())
	case "char":
		v := utils.ReadByte(r)
		return wrap(t, v)
	case "int", "SInt32", "UInt32":
		v := utils.Swap(utils.ReadInt(r))
		logrus.Debugf("int %x", v)
		return wrap(t, v)
	case "float":
		v := utils.ReadFloat32(r, binary.LittleEndian)
		logrus.Debugf("float:%x", v)
		return wrap(t, v)
	case "Array":
		return parseArray(r, t, false)
	case "map", "vector":
		return parseCollection(r, t, false)
	default:
		logrus.Debugf("%s (default)", t.Type)
		return parseObject(r, t)
	}
}

func parseArray(r io.Reader, t TypeField, align bool) UnityObject {
	size := utils.Swap(utils.ReadInt(r))
	logrus.Debugf("[]%s size: 0x%x", t.Children[1].Type, size)

	array := make([]UnityObject, size, size)
	for i := uint32(0); i < size; i++ {
		v := parseValue(r, *t.Children[1])
		array[i] = v
		s, _ := json.Marshal(v)
		logrus.Debugf("Array[%d] = %s", i, s)
	}
	if align {
		for i := 0; i < 4; i++ {
			if (int(t.Children[1].Size)*int(size)+i)%4 == 0 {
				break
			}
			_ = utils.ReadByte(r)
		}
	}

	return UnityObject{t: &t, Children: array}
}

func parseCollection(r io.Reader, t TypeField, align bool) UnityObject {
	logrus.Debugf("Collection(%s)", t.Type)
	return parseArray(r, *t.Children[0], align)
}
