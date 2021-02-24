/*
 * @Author: hsm
 * @Date: 2021-02-24 14:41:19
 * @Last Modified by: mikey.zhaopeng
 * @Last Modified time: 2021-02-24 17:26:15
 */

package jpath

import (
	"fmt"
	"log"
	"strconv"
)

// Object 对象
type object struct {
	Start int
	End   int
}

type indexes struct {
	Start int
	End   int // End == 0 时 表达 [1] 单索引形式
}

// ConditionHandler 条件控制处理
type ConditionHandler func(cxt *Context) bool

func (o *object) Get(content []rune) []rune {
	return content[o.Start:o.End]
}

func (o *object) GetString(content []rune) string {
	return string(content[o.Start:o.End])
}

// Path 路径
type Path struct {
	Index  indexes // 整个路径解析的范围
	Target object  // 查找的对象
	Depth  int     // 查找的深度

	Type int // 1 - 2. find all or one target.(move next)  -1. back (move prev) 0. root(代表当前)

	Condition ConditionHandler // 该路径是否符合条件标准

	Prev *Path
	Next *Path
}

func skipSpace(content []rune, i *int) {
	for *i < len(content) && content[*i] == ' ' {
		*i++
	}
}

//go:generate stringer -type=nexttype
type nexttype int

const (
	// nIndexes 范围 [] [:]所有 [1] [1:] [:2] [1:2]
	nIndexes nexttype = 1
	// nCondition 条件 ()
	nCondition nexttype = 2
	// nDepth 深度
	nDepth nexttype = 3
	// nNextPath 下个路径 /
	nNextPath nexttype = 0
)

func getTarget(content []rune, i *int) (head *object, nt nexttype) {

	skipSpace(content, i)
	n := *i

	head = &object{Start: n}
	defer func() {
		head.End = n
		*i = n
	}()

	for ; n < len(content); n++ {
		c := content[n]
		switch c {
		case '[': // 表示范围
			nt = nIndexes
			return
		case '(': // 条件
			nt = nCondition
			return
		case '<':
			nt = nDepth
			return
		case ' ':
			nn := n + 1
			for ; nn < len(content); nn++ {
				c := content[nn]
				switch c {
				case ' ':
					continue
				case '[': // 表示范围
					nt = nIndexes
					return
				case '<':
					nt = nDepth
					return
				case '(': // 条件
					nt = nCondition
					return
				case '/':
					nt = nNextPath
					return
				default:
					panic(fmt.Errorf("exp error"))
				}
			}
			return
		case '/': // 结束这段路径解析
			nt = nNextPath
			return
		default:
			continue
		}
	}

	return
}

func getIndexes(content []rune, i *int) (idxs *indexes, nt nexttype) {

	n := *i
	n++
	defer func() {
		*i = n
	}()

	// nIndexes 范围 [] [:]所有 [:2]  [1] [1:]  [1:2]
	idxs = &indexes{}
	c := content[n]
	if c == ']' { // []
		idxs.End = -1
		return
	} else if c == ':' {
		n++
		if content[n] == ']' { // [:]
			idxs.End = -1
			return
		}
		// [:2]

		// strconv.Atoi()
		var estr []rune
		for ; n < len(content); n++ {
			c = content[n]
			if c == ']' {
				break
			}
			estr = append(estr, c)
		}

		end, err := strconv.Atoi(string(estr))
		if err != nil {
			log.Panic(err, string(estr))
		}
		idxs.End = end
		return
	}

	for ; n < len(content); n++ {
		c := content[n]
		var sstr []rune
		for ; n < len(content); n++ {
			c = content[n]
			switch c {
			case ']':
				start, err := strconv.Atoi(string(sstr))
				if err != nil {
					log.Panic(err, string(sstr))
				}
				idxs.Start = start
				return

			case ':':
				start, err := strconv.Atoi(string(sstr))
				if err != nil {
					log.Panic(err, string(sstr))
				}
				idxs.Start = start

				n++
				if content[n] == ']' { // [:]
					idxs.End = -1
					return
				}

				var estr []rune
				for ; n < len(content); n++ {
					c = content[n]
					if c == ']' {
						break
					}
					estr = append(estr, c)
				}

				end, err := strconv.Atoi(string(estr))
				if err != nil {
					log.Panic(err, string(estr))
				}
				idxs.End = end
				return
			default:
				sstr = append(sstr, c)
				// panic(fmt.Sprintf("error format. %b", c))
			}

		}
	}

	return
}

func getDepth(content []rune, i *int) (depth int) {
	var err error

	n := *i
	n++
	defer func() {
		*i = n
	}()

	skipSpace(content, i)

	if content[n] == '>' {
		return -1
	}

	var depthstr []rune
	defer func() {
		depth, err = strconv.Atoi(string(depthstr))
		if err != nil {
			log.Panic(err, string(depthstr))
		}
	}()

	for ; n < len(content); n++ {
		c := content[n]
		if c == '>' {
			return
		}
		depthstr = append(depthstr, c)
	}
	return
}

func getCondition(content []rune, i *int) (condition string) {
	// var err error

	n := *i
	n++
	defer func() {
		*i = n
	}()

	skipSpace(content, i)

	if content[n] == ')' {
		return ""
	}

	var conditionstr []rune
	defer func() {
		condition = string(conditionstr)
	}()

	for ; n < len(content); n++ {
		c := content[n]
		if c == ')' {
			return
		}
		conditionstr = append(conditionstr, c)
	}
	return
}

func headHandler(src []rune) (content []rune) {

	defer func() {
		content = append(content, ' ')
	}()

	for i := 0; i < len(src); i++ {
		c := src[i]
		switch c {
		case ' ':
			continue
		case '/':
			content = append(content)
			content = append(content, src[i:]...)
			return
		case '(':
			content = append(content, '/', '.')
			content = append(content, src[i:]...)
			return
		case '<':
			content = append(content, '/', '.')
			content = append(content, src[i:]...)
			return
		case '[':
			content = append(content, '/', '.')
			content = append(content, src[i:]...)
			return
		default:
			content = append(content, '/')
			content = append(content, src[i:]...)
			return
		}
	}

	return
}

// Parse 解析操作路径
func Parse(src []rune) (result *Path) {

	// content := headHandler(src)

	// result = &Path{Type: 0}
	// cur := result

	// var i = 0
	// for i < len(content) {
	// 	target, nt := getTarget(content, &i)
	// 	switch nt {
	// 	case nIndexes:
	// 		indexes, nnt := getIndexes(content, &i)

	// 	case nCondition:
	// 	case nNextPath:
	// 	}
	// }

	return
}
