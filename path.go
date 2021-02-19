package jpath

// Path 路径
type Path struct {
	Type int // 1. find all or one target.(move next) 2. back (move prev)

	Target []byte // 查找的对象

	Condition func(cur []byte) bool

	Prev *Path
	Next *Path
}
