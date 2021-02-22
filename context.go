package jpath

type Target struct {
}

// Context Condition 上下文
type Context struct {
	Content []rune
	Current *Path
}

// GetTarget 获取目标对象内容
func (cxt *Context) GetTarget() *Target {
	return nil
}
