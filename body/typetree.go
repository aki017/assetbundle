package body

import "github.com/Sirupsen/logrus"

// TypeTree is typetree
type TypeTree struct {
	Version      int32
	UnityVersion string
	Fields       []*TypeField
}

func (t TypeTree) Get(id ClassID) *TypeField {
	for _, f := range t.Fields {
		if f.ID == id {
			return f
		}
	}
	logrus.Fatal("NONE")
	return nil
}
