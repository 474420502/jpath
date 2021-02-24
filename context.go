/*
 * @Author: hsm
 * @Date: 2021-02-24 14:42:51
 * @Last Modified by: mikey.zhaopeng
 * @Last Modified time: 2021-02-24 14:43:14
 */

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
