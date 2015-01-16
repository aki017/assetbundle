package header

import "fmt"

// Format is Enum
//   1 in Unity 1 to 2.5
//   2 in Unity 2.6 to 3.4
//   3 in Unity 3.5 and 4
type Format int32

func (f Format) String() string {
	switch f {
	case 1:
		return "Unity 1 to 2.5"
	case 2:
		return "Unity 2.6 to 3.4"
	case 3:
		return "Unity 3.5 and 4"
	default:
		return fmt.Sprintf("Format(%d)", int32(f))
	}
}
