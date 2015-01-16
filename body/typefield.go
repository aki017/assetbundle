package body

import "fmt"

// TypeField is typefield
type TypeField struct {
	Children  []*TypeField
	ID        ClassID
	Type      string
	Name      string
	Size      int32
	Index     int32
	ArrayFlag int32
	Flags1    int32
	Flags2    int32
}

func (t TypeField) Parse(d []byte) UnityObject {
	return ParseObject(d, t)
}

func (t TypeField) String() string {
	return fmt.Sprintf("%s %s", t.Type, t.Name)
}
