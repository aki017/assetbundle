package body

import (
	"fmt"
	"strings"
)

type UnityObject struct {
	Children []UnityObject
	data     interface{}
	t        *TypeField
}

func (u UnityObject) String() string {
	c := make([]string, len(u.Children))
	for i, o := range u.Children {
		c[i] = o.String()
		c[i] = "  " + strings.Join(strings.Split(c[i], "\n"), "\n  ")
	}
	if len(u.Children) == 0 {
		return fmt.Sprintf(`%s = %s`, u.t.Name, u.data)
	} else {
		return fmt.Sprintf(`%s = [
%s
]`, u.t.Name, strings.Join(c, "\n"))
	}
}
