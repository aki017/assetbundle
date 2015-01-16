package body

// Objects is array
type Objects struct {
	List []*Object
}

// Object is obj
type Object struct {
	ID       int32
	Offset   int32
	Length   int32
	ClassID1 ClassID
	ClassID2 ClassID
}
